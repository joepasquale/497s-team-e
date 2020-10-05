# 497s-team-e: LinkUp
GitHub repository for Team E's Scalable Web Systems Project. Github will be used as the Workflow measure and controller for this project, with Scrumboard on Github to trask essential task list. Calls will be asynchronously used as needed.  

This repo is the working directory for LinkUp a new group scheduler application that allows for democratic time management systems combined with effective scheduling algorithms to efficiently help groups of people plan schedules around their lives. Simply have endpoint users add their availibility into the system and LinkUp is able to choose effective time slots and suggest locations they can meet.

# a) Routing:
For this repo we are using Nginx in order to control routing between front-end and back-end. This is the industry standard and proper routing is essential for scalability. 
# b) REST API:
We are developing CRUD tasks using Flask to interact with the data structure. This is because of the familiarity our team already has with it. Also Flask has a wide knowledge base in industry.
# c) Connections:
To connect routing to our CRUD tasks we are using uWSGI as the simplest way to connect the two.
# d) Database
MongoDB is a NoSQL database for storing data for our web application. This DB is easy for our team to use and handles a large number of requests very well. MongoDB is useful for scaling. Data will be passed around in standardized JSON formats.
# e) Back-End
For now our back-end applications are used with Python, Go, and Javascript.
# f) Front-End
Our Front-End UI is designed and implemented with HTML+CSS+JS with utilization of ReactJS
# g) Microservice Architecture
For scalability purposes we are using Docker containers for different parts of our project and we will pull down Dockerfiles to build the images of the containers to satisfy requirements.
