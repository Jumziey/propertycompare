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
> property cost &ltprice> &ltoperating costs monthly> &ltproperty insurance monthly> &ltcurrent mortgage deed>
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
```
> property add &ltstreet address> &ltoperating costs monthly> &ltproperty insurance monthly> &ltcurrent mortgage deed>
+----------------------------------------------+--------+
| ADDED: Address -------------------------------------- |
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
### property delete
```
> property delete &ltaddress>
+----------------------------------------------+--------+
| DELETED: Address ------------------------------------ |
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
### property show
```
> property show &ltaddress>
+----------------------------------------------+--------+
| Address --------------------------------------------- |
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
### property list
``` 
> property list 
+----------------------------------------------+--------+
| Address 1-------------------------------------------- |
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

+----------------------------------------------+--------+
| Address 2-------------------------------------------- |
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
...
...
...

```
### property list --only-address
```
+----------------------------------------------+--------+
| Address 1-------------------------------------------- |
+----------------------------------------------+--------+
+----------------------------------------------+--------+
| Address 2-------------------------------------------- |
+----------------------------------------------+--------+
...
...
...

```
### property diff <address> <address>

## database
Want online syncing, looking at amazon dynamodb for add/remove/list directly. Do want to start with simple file list or similar to start with. Useful option to have and a great way to start.
