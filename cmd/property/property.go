package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

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

	taxProperty := propertycost.PropertyTax{
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

	fmt.Println(rentRebate)
	fmt.Println(taxMortgageDeed)
	fmt.Println(taxTitleDeed)
	fmt.Println(taxProperty)
	fmt.Println(mortgage)

	rootCmd := &cobra.Command{Use: "property"}
	rootCmd.AddCommand(&cobra.Command{
		Use:     "cost <price> <operating costs monthly> <property insurance monthly> <current mortgage deed>",
		Short:   "Calculates the monthly cost for the property",
		Long:    "It's calculating using all the rules relevant for small housing in Sweden",
		Version: "0.0.1",
		Args:    cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			price, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				log.Fatalw("Can't convert <price> to float", "error", err)
			}
			// operatingCostMonthly, err := strconv.ParseFloat(args[1], 64)
			// if err != nil {
			// 	log.Fatalw("Can't convert <operating cost monthly> to float", "error", err)
			// }
			// propertyInsuranceMonthly, err := strconv.ParseFloat(args[2], 64)
			// if err != nil {
			// 	log.Fatalw("Can't convert <propertyInsuranceMonthly> to float", "error", err)
			// }
			mortgageDeedCurrent, err := strconv.ParseFloat(args[3], 64)
			if err != nil {
				log.Fatalw("Can't convert <current mortgage deed> to float", "error", err)
			}
			extracost := propertycost.ExtraAtPurchase(price, mortgageDeedCurrent, taxMortgageDeed, taxTitleDeed)
			fmt.Println("Extra cost for this property: ", extracost)
		},
	})

	if err := rootCmd.Execute(); err != nil {
		log.Fatalw("cant execute rootCommand", "error", err)
	}
}
