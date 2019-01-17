## date-filler-api
This is a very simple microservice that returns the dates in between two given dates. The response includes some additional information about the dates such as `isoWeek` and stuff. If you pass a `unique` key to the `GET` call, it also returns a unique filling for the given key, meaning you can get continuous unique `isoWeek`s instead of days.

Sample response:
```json
[
  {
    "date": "2013-09-30",
    "year": 2013,
    "month": 9,
    "week": 40,
    "yearWeek": "201340",
    "day": 30,
    "isoYear": 2013,
    "isoWeek": 40,
    "isoYearIsoWeek": "201340",
    "weekday": 1,
    "yearday": 273
  },
  {
    "date": "2013-10-01",
    "year": 2013,
    "month": 10,
    "week": 40,
    "yearWeek": "201340",
    "day": 1,
    "isoYear": 2013,
    "isoWeek": 40,
    "isoYearIsoWeek": "201340",
    "weekday": 2,
    "yearday": 274
  }
]
```

## Fields
|Field|Type|Description|
|---|---|---|
|`date` | string | The exact date in `Y-m-d` format.|
|`year` | integer | Year of the date.|
|`month` | integer | Month index of the date, starting from 1.|
|`week` | integer | Week of the date, calculated by weeks starting on Sunday. |
|`yearWeek` | string | Year and week combined in format of `YYYYMM`. |
|`day` | integer | Day of the month.|
|`isoYear` | integer | ISO 8601 year of the given date.|
|`isoWeek` | integer | ISO 8601 week of the given date.|
|`isoYearIsoWeek` | string | ISO Year and ISO week combined in format of `YYYYMM`.|
|`weekday` | integer | The day index of the week, starting from 0 for Sunday. |
|`yearday` | integer | The day of the year. |

## Build
Simply run:
```bash
go build
```

Then you can run your executable directly:
```bash
./date-filler-api
```

## License
This project is under MIT license.
