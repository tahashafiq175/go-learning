package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	tommorrow := now.Add(time.Hour * 24)
    dayAfterTommorrow := now.Add(time.Hour * 24).Add(time.Hour * 24)

	fmt.Println(now)
	fmt.Println(tommorrow)
	fmt.Println(dayAfterTommorrow)

	fmt.Println("Current Hour:", now.Hour())
	fmt.Println("Current Minute:", now.Minute())
	fmt.Println("Current Second:", now.Second())
	fmt.Println("Current Year:", now.Year())
	fmt.Println("Current Month:", now.Month())
	fmt.Println("Current Day:", now.Day())
	fmt.Println("Current Location:", now.Location())

	birthday:=time.Date(
		2005,time.April,10,12,0,0,0,time.UTC,
	)
	fmt.Println("Birthday:", birthday)

	timeDifference := now.Sub(birthday)
	fmt.Println("Time Difference:", timeDifference)
	fmt.Println(timeDifference.Hours(), "Hours")

	//time Duration
	duration := time.Duration(2 * time.Hour)
	fmt.Println("Duration:", duration)


	//AddDate()

	newDate:= now.AddDate(1, 2, 3)
	fmt.Println("New Date after adding 1 year, 2 months and 3 days:", newDate)
	//simply AddDate() method is used to add or subtract years, months, and days from a given date. It takes three parameters: years, months, and days. You can pass negative values to subtract years, months, or days.
	newDate2:= now.AddDate(-1, -2, -3)
	fmt.Println("New Date after subtracting 1 year, 2 months and 3 days:", newDate2)

    endTime:=time.Now()

	duration2:=endTime.Sub(now)
	fmt.Println("Duration between start and end time:", duration2)		
	//The above code will print the current date and time in the default format. It will also print the date and time for tomorrow and the day after tomorrow. Additionally, it will print the current hour, minute, second, year, month, day, and location. It will also print the birthday date and the time difference between now and the birthday. Finally, it will print the duration between the start and end time.


	//

	// start:=time.Now()
	
	// time.Sleep(2* time.Second)

	// end:=time.Now()

	// diff:=end.Sub(start)


	// fmt.Println("it is the difference between " , diff)


	// //time.Since

	// fmt.Println(time.Since(start))


	// //time.Until

	// untilStart:=time.Now().Add(5*time.Second)
	// fmt.Println("Time until 5 seconds from now:", time.Until(untilStart))

	//the above code will print the current date and time in the default format. It will also print the date and time for tomorrow and the day after tomorrow. Additionally, it will print the current hour, minute, second, year, month, day, and location. It will also print the birthday date and the time difference between now and the birthday. Finally, it will print the duration between the start and end time.

    // Comparing times
//Before

	// if start.Before(end) {
	// 	fmt.Println("Start is before end")
	// }	

	//the purpose of Before is for comparing two time values and checking if the first time value is before the second time value. It returns a boolean value indicating whether the first time value is before the second time value or not.	



	start:=time.Date(2026,time.October,14,5,12,0,0,time.UTC)

	end:=time.Now()
	if(start.After(end)){
		fmt.Println("Deadline has not passed yet")
	}else{
		fmt.Println("Deadline has passed")
	}


// 	if t1.Equal(t2) {
//     fmt.Println("Same moment")
// }

	
// // The purpose of Equal is for comparing two time values and checking if they represent the same moment in time. It returns a boolean value indicating whether the two time values are equal or not.


// //different ways to declare time

// Ask Go for the current date and time, then save it in newNow.
// newNow := time.Now()
//
// Make a date that looks like year-month-day, for example 2026-10-14.
// Go uses this special example date to show where each part should go:
// 2006 is the year, 01 is the month, and 02 is the day.
// formattedDate := newNow.Format("2006-01-02")
// fmt.Println("Today's date is:", formattedDate)
//
// Make a time that looks like hour:minute:second, for example 17:45:17.
// 15 means the hour, 04 means the minute, and 05 means the second.
// timeText := newNow.Format("15:04:05")
// fmt.Println("The time is:", timeText)
//
// Make the date and time together, then print the result.
// The space between the date and time keeps them easy to read.
	formattedTimeDate := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(formattedTimeDate)
}


// Output:above code will print the current date and time in the default format.
// such that the main syntax is time.Time()  
// like time.now() will return the current local time and date in the default format.
// time.Now() returns the current local time and date in the default format.
// time.Hour() returns the current hour of the day in the default format.
// time.Minute() returns the current minute of the hour in the default format.
// time.Second() returns the current second of the minute in the default format.
// time.Add() returns the current time and date in the default format.
// time.AddDays() returns the current time and date in the default format.
// time.AddHours() returns the current time and date in the default format.
// time.AddMinutes() returns the current time and date in the default format. and so on.

