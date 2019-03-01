# NSQ PLugin
> Send specify message to NSQ topic

## How To Use?

"topic" and "message" must exist in the Ext attributes.

"topic" means dist nsq topic name. Also "message" means the transfer data. In the v0.1.0, the message only support string type.

## About the custom formate

All the custom date format string has a pre-string, "formate:<value kind>". 

|Value Kind |Format  | Value |
| ------ | ------ | ------ | 
| Date | ${Placeholder} | YYYY:year, MM:month, DD:Day, hh:hour, mm:minute, ss:second |



## Example

* Custom Data Format 
```
message=formate:${YYYYMMDD}
```


