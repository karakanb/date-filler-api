package main

import (
	"testing"
	"time"
)

func TestSundayWeek(t *testing.T) {
	tables := []struct {
		date         string
		expectedWeek int
	}{
		{"2018-12-31", 53},
		{"2018-12-15", 50},
		{"2018-11-18", 47},
		{"2018-11-03", 44},
		{"2018-10-12", 41},
		{"2018-10-05", 40},
		{"2018-09-08", 36},
		{"2018-08-29", 35},
		{"2018-08-08", 32},
		{"2018-07-30", 31},
		{"2018-07-07", 27},
		{"2018-07-05", 27},
		{"2018-07-04", 27},
		{"2018-06-12", 24},
		{"2018-06-02", 22},
		{"2018-05-04", 18},
		{"2018-04-29", 18},
		{"2018-04-09", 15},
		{"2018-03-26", 13},
		{"2018-03-11", 11},
		{"2018-02-22", 8},
		{"2018-02-20", 8},
		{"2018-01-23", 4},
		{"2017-12-27", 52},
		{"2017-12-02", 48},
		{"2017-11-22", 47},
		{"2017-11-18", 46},
		{"2017-10-22", 43},
		{"2017-10-02", 40},
		{"2017-09-04", 36},
		{"2017-08-22", 34},
		{"2017-08-09", 32},
		{"2017-08-02", 31},
		{"2017-08-01", 31},
		{"2017-07-14", 28},
		{"2017-06-22", 25},
		{"2017-05-27", 21},
		{"2017-05-05", 18},
		{"2017-05-01", 18},
		{"2017-04-02", 14},
		{"2017-03-23", 12},
		{"2017-02-26", 9},
		{"2017-02-06", 6},
		{"2017-01-09", 2},
		{"2017-01-05", 1},
		{"2016-12-09", 50},
		{"2016-11-26", 48},
		{"2016-11-06", 46},
		{"2016-10-29", 44},
		{"2016-10-19", 43},
		{"2016-10-09", 42},
		{"2016-10-06", 41},
		{"2016-09-11", 38},
		{"2016-08-30", 36},
		{"2016-08-04", 32},
		{"2016-07-23", 30},
		{"2016-07-10", 29},
		{"2016-06-11", 24},
		{"2016-05-23", 22},
		{"2016-05-06", 19},
		{"2016-05-05", 19},
		{"2016-04-08", 15},
		{"2016-03-14", 12},
		{"2016-02-19", 8},
		{"2016-01-24", 5},
		{"2016-01-11", 3},
		{"2016-01-04", 2},
		{"2015-12-25", 52},
		{"2015-11-26", 48},
		{"2015-11-16", 47},
		{"2015-10-23", 43},
		{"2015-10-05", 41},
		{"2015-09-09", 37},
		{"2015-09-02", 36},
		{"2015-08-31", 36},
		{"2015-08-01", 31},
		{"2015-07-18", 29},
		{"2015-07-01", 27},
		{"2015-06-10", 24},
		{"2015-05-24", 22},
		{"2015-05-19", 21},
		{"2015-05-12", 20},
		{"2015-04-26", 18},
		{"2015-04-21", 17},
		{"2015-04-08", 15},
		{"2015-03-31", 14},
		{"2015-03-29", 14},
		{"2015-03-06", 10},
		{"2015-02-20", 8},
		{"2015-01-26", 5},
		{"2015-01-23", 4},
		{"2014-12-24", 52},
		{"2014-12-23", 52},
		{"2014-12-05", 49},
		{"2014-11-07", 45},
		{"2014-10-28", 44},
		{"2014-10-25", 43},
		{"2014-09-30", 40},
		{"2014-09-23", 39},
		{"2014-09-04", 36},
		{"2014-08-10", 33},
		{"2014-07-17", 29},
		{"2014-07-11", 28},
		{"2014-06-16", 25},
		{"2014-05-24", 21},
		{"2014-05-13", 20},
		{"2014-04-13", 16},
		{"2014-03-22", 12},
		{"2014-03-02", 10},
		{"2014-02-04", 6},
		{"2014-01-09", 2},
		{"2014-01-05", 2},
		{"2013-12-23", 52},
		{"2013-12-14", 50},
		{"2013-12-04", 49},
		{"2013-11-24", 48},
		{"2013-11-21", 47},
		{"2013-10-24", 43},
		{"2013-10-22", 43},
		{"2013-09-30", 40},
		{"2013-09-04", 36},
		{"2013-08-24", 34},
		{"2013-08-20", 34},
		{"2013-07-29", 31},
		{"2013-07-23", 30},
		{"2013-07-15", 29},
		{"2013-07-07", 28},
		{"2013-07-05", 27},
		{"2013-06-25", 26},
		{"2013-05-30", 22},
		{"2013-05-09", 19},
		{"2013-04-28", 18},
		{"2013-04-08", 15},
		{"2013-03-28", 13},
		{"2013-03-03", 10},
		{"2013-02-26", 9},
		{"2013-02-16", 7},
		{"2013-01-24", 4},
		{"2012-12-28", 52},
		{"2012-12-26", 52},
		{"2012-12-01", 48},
		{"2012-11-03", 44},
		{"2012-11-02", 44},
		{"2012-10-22", 43},
		{"2012-10-17", 42},
		{"2012-10-08", 41},
		{"2012-10-02", 40},
		{"2012-09-30", 40},
		{"2012-09-02", 36},
		{"2012-08-16", 33},
		{"2012-07-31", 31},
		{"2012-07-07", 27},
		{"2012-06-08", 23},
		{"2012-05-26", 21},
		{"2012-05-24", 21},
		{"2012-04-30", 18},
		{"2012-04-22", 17},
		{"2012-04-19", 16},
		{"2012-04-07", 14},
		{"2012-03-24", 12},
		{"2012-03-09", 10},
		{"2012-02-19", 8},
		{"2012-02-02", 5},
		{"2012-01-14", 2},
		{"2011-12-20", 52},
		{"2011-12-11", 51},
		{"2011-12-07", 50},
		{"2011-11-14", 47},
		{"2011-10-25", 44},
		{"2011-10-03", 41},
		{"2011-09-29", 40},
		{"2011-08-30", 36},
		{"2011-08-14", 34},
		{"2011-07-22", 30},
		{"2011-07-06", 28},
		{"2011-06-26", 27},
		{"2011-06-20", 26},
		{"2011-06-01", 23},
		{"2011-05-16", 21},
		{"2011-05-12", 20},
		{"2011-04-23", 17},
		{"2011-03-24", 13},
		{"2011-03-18", 12},
		{"2011-03-06", 11},
		{"2011-02-19", 8},
		{"2011-02-16", 8},
		{"2011-02-07", 7},
		{"2011-01-31", 6},
		{"2011-01-08", 2},
		{"2010-12-09", 50},
		{"2010-11-30", 49},
		{"2010-11-15", 47},
		{"2010-10-18", 43},
		{"2010-09-29", 40},
		{"2010-09-11", 37},
		{"2010-08-22", 35},
		{"2010-08-18", 34},
		{"2010-08-09", 33},
		{"2010-07-21", 30},
		{"2010-06-28", 27},
		{"2010-06-19", 25},
		{"2010-06-02", 23},
		{"2010-05-23", 22},
		{"2010-05-13", 20},
		{"2010-04-18", 17},
		{"2010-04-03", 14},
		{"2010-03-12", 11},
		{"2010-02-23", 9},
		{"2010-02-05", 6},
		{"2010-01-20", 4},
		{"2010-01-08", 2},
		{"2009-12-17", 51},
		{"2009-11-17", 47},
		{"2009-11-09", 46},
		{"2009-10-10", 41},
		{"2009-09-12", 37},
		{"2009-08-29", 35},
		{"2009-08-09", 33},
		{"2009-08-08", 32},
		{"2009-07-10", 28},
		{"2009-07-02", 27},
		{"2009-06-08", 24},
		{"2009-05-23", 21},
		{"2009-05-13", 20},
		{"2009-04-20", 17},
		{"2009-03-28", 13},
		{"2009-03-05", 10},
		{"2009-02-13", 7},
		{"2009-02-03", 6},
		{"2009-01-23", 4},
		{"2009-01-12", 3},
		{"2009-01-01", 1},
		{"2008-12-25", 52},
		{"2008-12-15", 51},
		{"2008-12-11", 50},
		{"2008-12-03", 49},
		{"2008-11-12", 46},
		{"2008-11-10", 46},
		{"2008-10-18", 42},
		{"2008-10-02", 40},
		{"2008-09-14", 38},
		{"2008-09-12", 37},
		{"2008-09-11", 37},
		{"2008-08-18", 34},
		{"2008-07-29", 31},
		{"2008-06-30", 27},
		{"2008-06-05", 23},
		{"2008-05-06", 19},
		{"2008-05-01", 18},
		{"2008-04-22", 17},
		{"2008-04-07", 15},
		{"2008-04-05", 14},
		{"2008-03-26", 13},
		{"2008-03-03", 10},
		{"2008-02-03", 6},
		{"2008-01-24", 4},
		{"2008-01-18", 3},
		{"2007-12-30", 53},
		{"2007-12-09", 50},
		{"2007-12-01", 48},
		{"2007-11-11", 46},
		{"2007-10-31", 44},
		{"2007-10-19", 42},
		{"2007-09-26", 39},
		{"2007-09-19", 38},
		{"2007-08-29", 35},
		{"2007-08-21", 34},
		{"2007-08-12", 33},
		{"2007-07-17", 29},
		{"2007-07-12", 28},
		{"2007-06-25", 26},
		{"2007-06-01", 22},
		{"2007-05-04", 18},
		{"2007-04-06", 14},
		{"2007-03-18", 12},
		{"2007-03-04", 10},
		{"2007-02-09", 6},
		{"2007-01-16", 3},
		{"2007-01-10", 2},
		{"2006-12-20", 51},
		{"2006-12-12", 50},
		{"2006-11-14", 46},
		{"2006-10-27", 43},
		{"2006-10-08", 41},
		{"2006-09-29", 39},
		{"2006-09-14", 37},
		{"2006-09-08", 36},
		{"2006-08-28", 35},
		{"2006-08-26", 34},
		{"2006-08-12", 32},
		{"2006-08-09", 32},
		{"2006-07-29", 30},
		{"2006-07-06", 27},
		{"2006-06-28", 26},
		{"2006-06-15", 24},
		{"2006-05-16", 20},
		{"2006-04-18", 16},
		{"2006-03-26", 13},
		{"2006-03-25", 12},
		{"2006-03-11", 10},
		{"2006-02-26", 9},
		{"2006-02-03", 5},
		{"2006-01-09", 2},
		{"2005-12-17", 51},
		{"2005-11-22", 48},
		{"2005-11-07", 46},
		{"2005-10-10", 42},
		{"2005-09-25", 40},
		{"2005-08-27", 35},
		{"2005-08-07", 33},
		{"2005-07-11", 29},
		{"2005-06-15", 25},
		{"2005-05-29", 23},
		{"2005-05-16", 21},
		{"2005-05-10", 20},
		{"2005-04-16", 16},
		{"2005-03-25", 13},
		{"2005-02-28", 10},
		{"2005-02-23", 9},
		{"2005-02-01", 6},
		{"2005-01-08", 2},
		{"2004-12-13", 51},
		{"2004-11-25", 48},
		{"2004-10-30", 44},
		{"2004-09-30", 40},
		{"2004-09-21", 39},
		{"2004-09-15", 38},
		{"2004-09-06", 37},
		{"2004-08-08", 33},
		{"2004-07-28", 31},
		{"2004-06-30", 27},
		{"2004-06-14", 25},
		{"2004-05-16", 21},
		{"2004-05-12", 20},
		{"2004-04-19", 17},
		{"2004-03-21", 13},
		{"2004-02-27", 9},
		{"2004-02-17", 8},
		{"2004-02-12", 7},
		{"2004-01-15", 3},
		{"2004-01-14", 3},
		{"2003-12-24", 52},
		{"2003-11-24", 48},
		{"2003-11-22", 47},
		{"2003-10-26", 44},
		{"2003-09-27", 39},
		{"2003-09-04", 36},
		{"2003-09-03", 36},
		{"2003-08-05", 32},
		{"2003-07-18", 29},
		{"2003-07-06", 28},
		{"2003-07-02", 27},
		{"2003-06-10", 24},
		{"2003-05-15", 20},
		{"2003-05-05", 19},
		{"2003-04-29", 18},
		{"2003-03-31", 14},
		{"2003-03-12", 11},
		{"2003-03-07", 10},
		{"2003-02-22", 8},
		{"2003-02-02", 6},
		{"2003-01-15", 3},
		{"2002-12-31", 53},
		{"2002-12-11", 50},
		{"2002-11-16", 46},
		{"2002-10-25", 43},
		{"2002-10-08", 41},
		{"2002-10-04", 40},
		{"2002-09-29", 40},
		{"2002-09-14", 37},
		{"2002-08-19", 34},
		{"2002-08-07", 32},
		{"2002-07-15", 29},
		{"2002-07-03", 27},
		{"2002-06-28", 26},
		{"2002-06-21", 25},
		{"2002-06-11", 24},
		{"2002-05-25", 21},
		{"2002-05-07", 19},
		{"2002-04-24", 17},
		{"2002-03-29", 13},
		{"2002-03-02", 9},
		{"2002-02-17", 8},
		{"2002-02-16", 7},
		{"2002-02-08", 6},
		{"2002-01-18", 3},
		{"2001-12-22", 51},
		{"2001-11-27", 48},
		{"2001-11-17", 46},
		{"2001-11-16", 46},
		{"2001-11-12", 46},
		{"2001-10-18", 42},
		{"2001-10-17", 42},
		{"2001-10-09", 41},
		{"2001-10-08", 41},
		{"2001-10-03", 40},
		{"2001-09-05", 36},
		{"2001-09-01", 35},
		{"2001-08-29", 35},
		{"2001-08-27", 35},
		{"2001-08-08", 32},
		{"2001-07-16", 29},
		{"2001-07-09", 28},
		{"2001-06-28", 26},
		{"2001-05-30", 22},
		{"2001-05-27", 22},
		{"2001-04-28", 17},
		{"2001-04-25", 17},
		{"2001-04-18", 16},
		{"2001-04-06", 14},
		{"2001-03-23", 12},
		{"2001-02-22", 8},
		{"2001-02-18", 8},
		{"2001-02-04", 6},
		{"2001-01-13", 2},
		{"2000-12-22", 52},
		{"2000-12-09", 50},
		{"2000-12-02", 49},
		{"2000-11-07", 46},
		{"2000-10-17", 43},
		{"2000-10-09", 42},
		{"2000-09-13", 38},
		{"2000-08-31", 36},
		{"2000-08-02", 32},
		{"2000-07-20", 30},
		{"2000-07-18", 30},
		{"2000-06-30", 27},
		{"2000-05-31", 23},
		{"2000-05-17", 21},
		{"2000-05-10", 20},
		{"2000-04-10", 16},
		{"2000-04-01", 14},
		{"2000-03-16", 12},
		{"2000-02-19", 8},
		{"2000-01-29", 5},
		{"2000-01-17", 4},
		{"1999-12-20", 52},
		{"1999-11-28", 49},
		{"1999-11-07", 46},
		{"1999-10-12", 42},
		{"1999-10-10", 42},
		{"1999-10-01", 40},
		{"1999-09-15", 38},
		{"1999-08-20", 34},
		{"1999-07-28", 31},
		{"1999-07-26", 31},
		{"1999-06-30", 27},
		{"1999-06-08", 24},
		{"1999-05-16", 21},
		{"1999-04-23", 17},
		{"1999-03-25", 13},
		{"1999-03-20", 12},
		{"1999-03-19", 12},
		{"1999-03-17", 12},
		{"1999-03-13", 11},
		{"1999-03-04", 10},
		{"1999-03-01", 10},
		{"1999-02-15", 8},
		{"1999-01-24", 5},
		{"1999-01-18", 4},
		{"1998-12-22", 52},
		{"1998-12-06", 50},
		{"1998-11-23", 48},
		{"1998-10-24", 43},
		{"1998-10-12", 42},
		{"1998-09-23", 39},
		{"1998-09-10", 37},
		{"1998-08-18", 34},
		{"1998-07-21", 30},
		{"1998-06-26", 26},
		{"1998-06-15", 25},
		{"1998-06-11", 24},
		{"1998-05-31", 23},
		{"1998-05-02", 18},
		{"1998-04-04", 14},
		{"1998-04-01", 14},
		{"1998-03-29", 14},
		{"1998-03-13", 11},
		{"1998-03-12", 11},
		{"1998-02-13", 7},
		{"1998-01-22", 4},
		{"1998-01-11", 3},
		{"1998-01-09", 2},
		{"1998-01-07", 2},
		{"1998-01-01", 1},
		{"1997-12-09", 50},
		{"1997-12-05", 49},
		{"1997-11-06", 45},
		{"1997-10-20", 43},
		{"1997-09-22", 39},
		{"1997-09-15", 38},
		{"1997-08-19", 34},
		{"1997-07-28", 31},
		{"1997-07-10", 28},
		{"1997-06-28", 26},
		{"1997-05-30", 22},
		{"1997-05-26", 22},
		{"1997-05-14", 20},
		{"1997-04-18", 16},
		{"1997-04-11", 15},
		{"1997-03-31", 14},
		{"1997-03-12", 11},
		{"1997-03-10", 11},
		{"1997-02-16", 8},
		{"1997-02-09", 7},
		{"1997-01-20", 4},
		{"1996-12-21", 51},
		{"1996-12-11", 50},
		{"1996-11-18", 47},
		{"1996-11-13", 46},
		{"1996-11-11", 46},
		{"1996-10-26", 43},
		{"1996-10-10", 41},
		{"1996-10-04", 40},
		{"1996-09-27", 39},
		{"1996-08-29", 35},
		{"1996-08-25", 35},
		{"1996-08-22", 34},
		{"1996-08-20", 34},
		{"1996-08-09", 32},
		{"1996-08-04", 32},
		{"1996-07-14", 29},
		{"1996-06-16", 25},
		{"1996-05-29", 22},
		{"1996-05-09", 19},
		{"1996-04-16", 16},
		{"1996-04-11", 15},
		{"1996-03-25", 13},
		{"1996-03-21", 12},
		{"1996-03-05", 10},
		{"1996-03-02", 9},
		{"1996-02-04", 6},
		{"1996-01-16", 3},
		{"1995-12-19", 51},
		{"1995-12-02", 48},
		{"1995-11-26", 48},
		{"1995-11-12", 46},
		{"1995-11-11", 45},
		{"1995-10-13", 41},
	}

	for _, table := range tables {

		t.Run(table.date, func(t *testing.T) {
			timeInstance, _ := time.Parse(dateFormat, table.date)
			calculatedWeek := sundayWeek(timeInstance)
			if calculatedWeek != table.expectedWeek {
				t.Errorf("sundayWeek was incorrect, got: %d, want: %d.", calculatedWeek, table.expectedWeek)
			}
		})
	}

}
