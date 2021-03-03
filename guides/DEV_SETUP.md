
Developer environment options
=================

We've made a few attempts at simplifying the way we run our microservices locally, but haven't converged on one way of doing things yet. A few options are explored in this guide. 

### iTerm2 setup
  - Any terminal can be used however there are advantages of using iTerm2 such as badges, profiles, arrangements, multiple panes, ability to send commands to all panes.
  - [Install iterm2](https://www.iterm2.com/)
  - [Setup the pane arrangements and profiles](https://blog.andrewray.me/how-to-create-custom-iterm2-window-arrangments/)
  - Optional: [Give each profile a badge](https://www.iterm2.com/documentation-badges.html)

### websysd
[websysd](https://github.com/ONSdigital/dp/tree/main/websysd) can accomplish some steps of getting local services running for you.

It contains functionality for installing some prerequisite services directly to your machine (rather than using dp-compose as the guide states), cloning most of the repositories and a manifest to be able to update a single file to configure all services prior to running them.

It is not maintained, and the run scripts will need to be updated to specify locations on your machine, so it is not an out of the box solution.

--------------


  ### Useful resources
* [Splash-page](https://github.com/ONSdigital/dp-setup/tree/develop/splash-page) (if you get a 404 error following this link, you need to be added to the ONSdigital organisation in GitHub), a list of digital publishing's repositories, environments, and platform-tools
* [dp-cli](https://github.com/ONSdigital/dp-cli) is a useful tool to grant access to environments and SSH into them when working offsite as well as other tasks
