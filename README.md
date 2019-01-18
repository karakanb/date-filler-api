## date-filler-api
This is a very simple microservice that returns the dates in between two given dates. The response includes some additional information about the dates such as `isoWeek` and stuff. If you pass a `unique` key to the `GET` call, it also returns a unique filling for the given key, meaning you can get continuous unique `isoWeek`s instead of days.

Sample call:
```curl
curl -X GET 'http://localhost:8080?start=2013-12-30&end=2014-01-01'
```

Sample response:
```json
[
    {
        "date": "2013-12-30",
        "year": 2013,
        "month": 12,
        "week": 53,
        "yearWeek": "201353",
        "day": 30,
        "isoYear": 2014,
        "isoWeek": 1,
        "isoYearIsoWeek": "201401",
        "weekday": 1,
        "yearday": 364
    },
    {
        "date": "2013-12-31",
        "year": 2013,
        "month": 12,
        "week": 53,
        "yearWeek": "201353",
        "day": 31,
        "isoYear": 2014,
        "isoWeek": 1,
        "isoYearIsoWeek": "201401",
        "weekday": 2,
        "yearday": 365
    },
    {
        "date": "2014-01-01",
        "year": 2014,
        "month": 1,
        "week": 1,
        "yearWeek": "201401",
        "day": 1,
        "isoYear": 2014,
        "isoWeek": 1,
        "isoYearIsoWeek": "201401",
        "weekday": 3,
        "yearday": 1
    }
]
```

## Usage
Just grab the Docker image:
```bash
docker run -p 8080:80 karakanb/date-filler-api:latest
```

The image runs on the port 80 of the container, so you can map it to wherever you want. Also, if you want to change the default port, you can pass `PORT` environment variable.
```bash
docker run -p 8080:8085 -e PORT=8085 karakanb/date-filler-api:latest
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

### Unique Fields
API allows you to get the values on a unique dimension, which would allow you to retrieve the first items of a given dimension. For example, if you ask `month` to be unique for a given period you will only get the first instance for each month.

``` curl
curl -X GET 'http://localhost:8080?start=2013-12-20&end=2014-03-03&unique=month'
``` 

gets a response like this:
```json
[
    {
        "date": "2013-12-20",
        "year": 2013,
        "month": 12,
        "week": 51,
        "yearWeek": "201351",
        "day": 20,
        "isoYear": 2013,
        "isoWeek": 51,
        "isoYearIsoWeek": "201351",
        "weekday": 5,
        "yearday": 354
    },
    {
        "date": "2014-01-01",
        "year": 2014,
        "month": 1,
        "week": 1,
        "yearWeek": "201401",
        "day": 1,
        "isoYear": 2014,
        "isoWeek": 1,
        "isoYearIsoWeek": "201401",
        "weekday": 3,
        "yearday": 1
    },
    {
        "date": "2014-02-01",
        "year": 2014,
        "month": 2,
        "week": 5,
        "yearWeek": "201405",
        "day": 1,
        "isoYear": 2014,
        "isoWeek": 5,
        "isoYearIsoWeek": "201405",
        "weekday": 6,
        "yearday": 32
    },
    {
        "date": "2014-03-01",
        "year": 2014,
        "month": 3,
        "week": 9,
        "yearWeek": "201409",
        "day": 1,
        "isoYear": 2014,
        "isoWeek": 9,
        "isoYearIsoWeek": "201409",
        "weekday": 6,
        "yearday": 60
    }
]
```

## Build
Build the project:
```bash
go build -o api
```

Then you can run your executable directly:
```bash
./api
```

## License
This project is under MIT license.
