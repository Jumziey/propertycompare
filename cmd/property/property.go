package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jumziey/propertycost/pkg/propertycost"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logConf := zap.NewDevelopmentConfig()
	logConf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := logConf.Build()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	log := logger.Sugar()

	flag.String("config", ".", "config file")
	flag.Parse()

	conf := viper.New()
	if err := conf.BindPFlags(flag.CommandLine); err != nil {
		log.Fatalw("can not parse CLI flags",
			"error", err,
		)
	}

	conf.SetConfigName("config")
	conf.AddConfigPath(conf.GetString("config"))
	conf.AddConfigPath(".")
	ex, err := os.Executable()
	if err != nil {
		log.Fatalw("can not find my own path", "error", err)
	}
	conf.AddConfigPath(filepath.Dir(ex))
	conf.SetConfigType("toml")
	if err := conf.ReadInConfig(); err != nil {
		log.Errorw("error reading config", "error", err)
	}

	rentRebate := propertycost.RentRebate{
		Limit:       conf.GetFloat64("rent-rebate.limit"),
		BeforeLimit: conf.GetFloat64("rent-rebate.before-limit"),
		AfterLimit:  conf.GetFloat64("rent-rebate.after-limit"),
	}
	taxMortgageDeed := conf.GetFloat64("tax.mortage-deed")
	taxTitleDeed := conf.GetFloat64("tax.title-deed")

	taxProperty := propertycost.TaxProperty{
		TaxationValuePercentageOfValue: conf.GetFloat64("tax.property.taxation-value-percentage-of-value"),
		Percent:                        conf.GetFloat64("tax.property.percent"),
		Roof:                           conf.GetFloat64("tax.property.roof"),
	}

	mortgage := propertycost.Mortgage{
		Rent:         conf.GetFloat64("mortgage.rent"),
		Amortization: conf.GetFloat64("mortgage.amortization"),
		DownPayment: propertycost.DownPayment{
			AmountInHand:       conf.GetFloat64("mortgage.down-payment.amount-in-hand"),
			RequiredPercentage: conf.GetFloat64("mortgage.down-payment.required-percentage"),
			Rent:               conf.GetFloat64("mortgage.down-payment.rent"),
			Amortization:       conf.GetFloat64("mortgage.down-payment.amortization"),
		},
	}

	cost := &cobra.Command{
		Use:     "cost <price> <operating costs yearly> <property insurance monthly> <current mortgage deed>",
		Short:   "Calculates the monthly cost for the property",
		Long:    "It's calculating using all the rules relevant for small housing in Sweden",
		Version: "0.0.1",
	}

	costHouse := &cobra.Command{
		Use:     "house <price> <operating costs yearly> <property insurance monthly> <current mortgage deed>",
		Short:   "Calculates the monthly cost for the property",
		Long:    "It's calculating using all the rules relevant for small housing in Sweden",
		Version: "0.0.1",
		Args:    cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			price, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				log.Fatalw("Can't convert <price> to float", "error", err)
			}
			operatingCostMonthly, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				log.Fatalw("Can't convert <operating cost monthly> to float", "error", err)
			}
			propertyInsuranceMonthly, err := strconv.ParseFloat(args[2], 64)
			if err != nil {
				log.Fatalw("Can't convert <propertyInsuranceMonthly> to float", "error", err)
			}
			mortgageDeedCurrent, err := strconv.ParseFloat(args[3], 64)
			if err != nil {
				log.Fatalw("Can't convert <current mortgage deed> to float", "error", err)
			}
			purchaseFees := propertycost.HousePurchaseFees(price, mortgageDeedCurrent, taxMortgageDeed, taxTitleDeed)

			realCostMonthly, amortizationMonthly, err := propertycost.HouseMonthly(price, operatingCostMonthly, mortgage, rentRebate, taxProperty, propertyInsuranceMonthly)
			if err != nil {
				log.Fatalw("can't calculate monthly costs", "error", err)
			}
			mRent, dpRent, err := propertycost.Rent(price, mortgage)
			if err != nil {
				log.Fatalw("can't calculate rent", "error", err)
			}
			rebate := propertycost.Rebate(mRent+dpRent, rentRebate)
			taxPropertyCost := propertycost.HouseTax(price, taxProperty)

			t := table.NewWriter()
			t.SetStyle(table.StyleLight)
			t.SetOutputMirror(os.Stdout)
			t.AppendRows([]table.Row{
				{"Total cash needed outside mortgage", fmt.Sprintf("%.1f", propertycost.RequiredDownPayment(price, mortgage.DownPayment)+purchaseFees)},
				{"Purchase fees", fmt.Sprintf("%.1f", purchaseFees)},
				{"Down payment required", fmt.Sprintf("%.1f", propertycost.RequiredDownPayment(price, mortgage.DownPayment))},
			})
			t.AppendSeparator()
			t.AppendRows([]table.Row{
				{"Monthly payment with rebate and tax", fmt.Sprintf("%.1f", realCostMonthly+amortizationMonthly)},
				{"Monthly payment without rebate and tax", fmt.Sprintf("%.1f", realCostMonthly+amortizationMonthly-taxPropertyCost/12+rebate/12)},
				{"Real cost monthly with rebate and tax", fmt.Sprintf("%.1f", realCostMonthly)},
				{"Amortization monthly", fmt.Sprintf("%.1f", amortizationMonthly)},
			})
			t.Render()
		},
	}
	cost.AddCommand(costHouse)

	costCondo := &cobra.Command{
		Use:     "condo <price> <operating costs yearly> <property insurance monthly>",
		Short:   "Calculates the monthly cost for the property",
		Long:    "It's calculating using all the rules relevant for condominium in Sweden",
		Version: "0.0.1",
		Args:    cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			price, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				log.Fatalw("Can't convert <price> to float", "error", err)
			}
			operatingCostMonthly, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				log.Fatalw("Can't convert <operating cost monthly> to float", "error", err)
			}
			propertyInsuranceMonthly, err := strconv.ParseFloat(args[2], 64)
			if err != nil {
				log.Fatalw("Can't convert <propertyInsuranceMonthly> to float", "error", err)
			}

			realCostMonthly, amortizationMonthly, err := propertycost.CondoMonthly(price, operatingCostMonthly, mortgage, rentRebate, propertyInsuranceMonthly)
			if err != nil {
				log.Fatalw("can't calculate monthly costs", "error", err)
			}
			mRent, dpRent, err := propertycost.Rent(price, mortgage)
			if err != nil {
				log.Fatalw("can't calculate rent", "error", err)
			}
			rebate := propertycost.Rebate(mRent+dpRent, rentRebate)

			t := table.NewWriter()
			t.SetStyle(table.StyleLight)
			t.SetOutputMirror(os.Stdout)
			t.AppendRows([]table.Row{
				{"Total cash needed outside mortgage", fmt.Sprintf("%.1f", propertycost.RequiredDownPayment(price, mortgage.DownPayment))},
				{"Down payment required", fmt.Sprintf("%.1f", propertycost.RequiredDownPayment(price, mortgage.DownPayment))},
			})
			t.AppendSeparator()
			t.AppendRows([]table.Row{
				{"Monthly payment with rebate", fmt.Sprintf("%.1f", realCostMonthly+amortizationMonthly)},
				{"Monthly payment without rebate", fmt.Sprintf("%.1f", realCostMonthly+amortizationMonthly+rebate/12)},
				{"Real cost monthly with rebate", fmt.Sprintf("%.1f", realCostMonthly)},
				{"Amortization monthly", fmt.Sprintf("%.1f", amortizationMonthly)},
			})
			t.Render()
		},
	}
	cost.AddCommand(costCondo)

	rootCmd := &cobra.Command{Use: "property"}
	rootCmd.AddCommand(cost)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalw("cant execute rootCommand", "error", err)
	}
}
