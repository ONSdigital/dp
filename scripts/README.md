Publish Metric Scripts
==========================

These scripts are designed to analyse collection data over time and provide insights into how and what we are publishing.

Publish-log-analysis provides breakdowns of content types and volumes week by week.

Publishing-performance-analysis looks at the start and end times of publish events compared with the amount of data going out, to measure 9.30 efficiency.

How to run the scripts
------

### Setup the data
Update the /publish-archive on the publishing mount to include all collection data you wish to analyse

* You can do stuff like `sudo cp -r 2018-07-1* /publish-archive/` to do it by 10 day batches but even that uses a fair amount of CPU and is quite slow. Much better to do it late on when publishing aren’t around

You can manually create a snapshot of the /publish-archive volume in AWS, and then create a new volume to mount to a new machine. This means running the script and chomping CPU won’t get in the way of live systems. Or you can run the scripts out of hours.

### Run the Go script to get the data

1. Checkout this repo locally.

2. Build the binary for whichever script you're running: 
`GOOS=linux GOARCH=amd64 go build -o metrics publish-log-analysis/main.go`

3. SCP the binary on to the box where your collections data is mounted, [these instructions may help](https://github.com/ONSdigital/dp-setup/tree/develop/scripts#ansible-scp)

4. SSH onto the box and check the executable file is in your home directory.

5. To run the script, you may want to use `screen` so you don’t lock up your shell or have to keep an active connection. The scripts may take in the region of an hour to run.

    The script outputs to STDOUT, so you'll likely want to redirect the output to a file.

    All of this can be accomplished with the steps:
    ```
    `screen`
    `./metrics > output.csv`
    CTRL+A+D
    ```

6. Copy the output to your local machine. Make sure it is called `output.csv` and that it’s in the same directory as the post-process.r script. (SCP is a good way to do this)


### Run the R script to generate the report

1. You’ll need to install R:
  * `Brew install r`
  * Or install from the [website](https://www.r-project.org/)

2. Open an R session by running the command `R`

3. In this session, install the dependencies:
`install.packages(c(“readr”))`

    You'll need to repeat this for every 'library' listed in the top of the R script you're running

4. Then run `source("publish-log-analysis/post-process.r”)`

5. And close the session with `q()`

By this point you should have an output in your local directory which you can send on to the publishing team member who requested the data.


