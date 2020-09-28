# Design Notes

## cmd:s
### `property cost`
Want a config with everything that wont change between properties, but for other circumstances. 
- [rent-rebate]
	- limit
	- before-limit
	- after-limit
- [property-tax]
	- taxation-value-percentage-of-value
	- percent
	- roof
- [mortgage]
	- rent
	- amortization
	- [mortage.down-payment]
		- amount-in-hand
		- required-percentage
		- rent
		- amortization
- mortage-deed-tax
- title-deed-tax
```
> property cost <price> <operating costs monthly> <property insurance monthly> <current mortgage deed>
+----------------------------------------------+--------+
| Total cash needed outside mortgage --------- | amount |
| One time cost at purchase                    | amount |
| down payment required                        | amount |
+----------------------------------------------+--------+
| Total monthly payment With Rebate And Tax    | amount |
| Total monthly payment Without Rebate And Tax | amount |
| Real cost monthly                            | amount |
| Amortization monthly                         | amount |
+----------------------------------------------+--------+
```
### property add
### property remove
### property list

## database
Want online syncing, looking at amazon dynamodb for add/remove/list directly. Do want to start with simple file list or similar to start with. Useful option to have and a great way to start.

## Then what?



