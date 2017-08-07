# jobber

## Time Strings

Field `time` specifies the schedule at which a job is run, in a manner similar to how cron jobs’ schedules are specified: with a space-separated list of up to six specifiers, each one of which constrains a different component of time (second, minute, hour, etc.). Schematically:

```
sec min hour month_day month week_day
```

*A job is scheduled thus*: it will run at any time that satisfies all of the specifiers sec, min, hour, and month and one of the specifiers `month_day` and `week_day`.

Each specifier can take one of the following forms: ("a", "b", "c", and "n" are placeholders for arbitrary numerals.)

| Specifier Form | What It Matches                                                                                                                       |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------|
| *              | Any value                                                                                                                             |
| a              | The value a                                                                                                                           |
| */n            | Every n-th value — e.g., */25 in the sec specifier would match 0, 25, and 50, whereas in the month specifier it would match 1 and 26. |
| a,b,c,...      | The values a, b, c, ...                                                                                                               |
| a-b            | Any of the values between and including a and b                                                                                       |

The specifiers have different permitted values for the placeholders in the specifier forms:

| Specifier | Values for “a”, “b”, and “c” | Values for "n" |
|-----------|------------------------------|----------------|
| sec       | 0 thru 59                    | 1 thru 59      |
| min       | 0 thru 59                    | 1 thru 59      |
| hour      | 0 thru 23                    | 1 thru 23      |
| month_day | 1 thru 31                    | 1 thru 30      |
| month     | 1 thru 12                    | 1 thru 11      |
| week_day  | 0 thru 6                     | 1 thru 5       |

NOTE: Because in a YAML document `*` has a special meaning at the beginning of an item, if your time string starts with `*` you must quote the whole string, thus: time: `'*/10'`.
