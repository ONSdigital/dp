#download csv and have it in working directory
library(tidyverse)
library(dplyr)
library(ggplot2)
library(scales)
library(reshape2)

df <- read.csv('/Users/iankent/dev/src/github.com/ONSdigital/dp/scripts/publishing-performance-analysis/publish-performance.csv')

#convert start date into days
df$start_date <- substr(df$PublishStartDate,1,10)
df$start_date <- as.Date(df$start_date, '%Y-%m-%d')

#remove NAs
df <- df[!is.na(df$start_date),]

#make sure data is numeric
df$FileCount <- as.numeric(df$FileCount)
df$FileSize <- as.numeric(df$FileSize) 
df$Duration <- as.numeric(df$Duration) 

df_scheduled <- df[!(df$Type == "scheduled"),]
df_manual <- df[!(df$Type == "manual"),]

#pivot to get values
pivot <- df %>%
  group_by(start_date) %>%
  summarise(sumFileCount = sum(FileCount), sumFileSize = sum(FileSize)/100000, maxDuration = max(Duration), aveDuration = mean(Duration), count = n()*1000)

pivot_scheduled <- df_scheduled %>%
  group_by(start_date) %>%
  summarise(sumFileCount = sum(FileCount), sumFileSize = sum(FileSize)/100000, maxDuration = max(Duration), aveDuration = mean(Duration), count = n()*1000)

pivot_manual <- df_manual %>%
  group_by(start_date) %>%
  summarise(sumFileCount = sum(FileCount), sumFileSize = sum(FileSize)/100000, maxDuration = max(Duration), aveDuration = mean(Duration), count = n()*1000)


#get count of entries

#entries <- df %>%
#  group_by(start_date) %>%
#  summarise(count = n())

#melted <- melt(pivot, id=c("start_date"))

# graph the pivot table by month
#ggplot(data = melted, aes(x=`strftime(start_date, "%Y/%m")`, y=value, colour=variable)) + geom_line(aes(group=variable)) + geom_smooth(method = "lm", aes(group=variable)) + xlab("Month") + ylab("Words") + ggtitle("Number of words by page type by month") + scale_y_continuous(label=comma)

pivot_x <- pivot[!(pivot$maxDuration > 100000),]
pivot_x_scheduled <- pivot_scheduled[!(pivot_scheduled$maxDuration > 100000),]
pivot_x_manual <- pivot_manual[!(pivot_manual$maxDuration > 100000),]

pivot_gathered <- pivot_x %>%
  gather(key = 'key', value = 'value', -start_date)
pivot_gathered_scheduled <- pivot_x_scheduled %>%
  gather(key = 'key', value = 'value', -start_date)
pivot_gathered_manual <- pivot_x_manual %>%
  gather(key = 'key', value = 'value', -start_date)

ggplot(pivot_gathered, aes(x = start_date, y = value, color = key)) + geom_line() + geom_smooth(method="lm", aes(group=key)) + scale_y_continuous(sec.axis = sec_axis(~./1000))
ggplot(pivot_gathered_scheduled, aes(x = start_date, y = value, color = key)) + geom_line() + geom_smooth(method="lm", aes(group=key)) + scale_y_continuous(sec.axis = sec_axis(~./1000))
ggplot(pivot_gathered_manual, aes(x = start_date, y = value, color = key)) + geom_line() + geom_smooth(method="lm", aes(group=key)) + scale_y_continuous(sec.axis = sec_axis(~./1000))
