![](react-frontend/src/assets/images/Coronavirus-Banner5.jpg)

# Covid-19 Web App (Go on AWS Kubernetes)

You may view the application by following the link below:

> [Virus Statistics](http://virus-go.s3-website-ap-southeast-2.amazonaws.com)


### Synopsis
The intention of this program is to allow users to connect via the web and retrieve daily statistical information on the 
**Covid-19** pandemic.
 
The application presents a summary of the disease activity, based on location. The information includes 
critical patients, daily deltas as well as the deaths caused by the disease over the recent period. A user may click upon a region to display a graph of the last month of activity.

Users may also sign-up for a newsletter which they will be sent via email, or they may opt to receive SMS notifications. 

### Technical Details
![](react-frontend/src/assets/images/Go_Logo.png)

This program was written as a Go (1.19) microservice, on an Amazon Elastic Kubernetes Service (EKS). This service is running as a 3 node cluster distributed over multiple AZ. 

The application uses Gin which provides a full-featured web framework. AWS DynamoDB provides the persistence tier, for storage of signed up users.

GitHub Actions loads compiles and pushes the code to Docker.

Test Driven Design (TDD) has been employed during the creation of this application. Frameworks such as httpMock and testify assisted during this phase (and ultimately by stubbing out interfaces).

The design pattern of *'separation of concerns'* is employed in this development. The layers are split up using the Go package mechanism, as shown below:

  - **main** -- Routes the request to the controller
  - **controller** -- This layer has direct access to the Web/HTTP layer. Its purpose is to mediate access to the service layer
  - **service** -- The service layer provides a boundary to the backend, exposed through a set of interfaces
  - **client** -- Controlled by the service layer. It provides access to 3rd party API's, message and event brokers
  - **custom_error** -- The set of application specific custom errors, that are typically used for communication between layers
  - **model** -- Used as the VO (i.e Value Objects) for communicating between layers and ultimately as JSON output to the web

The front end of the application was written using **React JS** ([React Hooks](https://reactjs.org/docs/hooks-intro.html))
together with front end HTML/CSS styling via [Bootstrap](https://getbootstrap.com).
 
 ### How to deploy the Application to AWS
 You must have the following to run this application:
 - Docker container
 - Make
 - An AWS account (available for free)
 - AWS CLI
 - node and NPM 
 
 Then perform the following:
 ```
 1. First, create an s3 bucket to store your static website:

 aws s3api create-bucket --bucket <BUCKET NAME> --acl public-read
 aws s3 website s3://<BUCKET NAME> --index-document index.html
 {
   "Version": "2012-10-17",
   "Statement": [
     {
       "Sid": "PublicReadForGetBucketObjects",
       "Effect": "Allow",
       "Principal": "*",
       "Action": "s3:GetObject",
       "Resource": "arn:aws:s3:::<BUCKET NAME>/*"
     }
   ]
 }
 aws s3api put-bucket-policy --bucket <BUCKET NAME> --policy file://policy_s3.json 
 
2. Edit the deploy.sh to refer to your bucket name
3. ./deploy.sh
4. Your site will be at: 
        http://<BUCKET NAME>.s3-website-ap-southeast-2.amazonaws.com
5. cd deployment/eks
6. make
7. cd ../react-frontend
8. make
9. cd ../go-src
10. make update        
 ```
 **Note**: It can be quite an **expensive** option to run an AWS EKS. This may be avoided by either running your application locally, using `make run` or inside a minikube installation.

### Further Readings

- [What’s so great about Go?](https://stackoverflow.blog/2020/11/02/go-golang-learn-fast-programming-languages/)
- [Amazon Elastic Kubernetes Service (EKS)](https://aws.amazon.com/eks/)
- [Gin Web Framework](https://gin-gonic.com)
- [GitHub Actions](https://docs.github.com/en/actions)
- [Stubbing Out Interfaces in Go](https://medium.com/stupid-gopher-tricks/stubbing-out-interfaces-in-go-e85afc200aa8)
- [ReactJS on AWS](https://viastudio.com/hosting-a-reactjs-app-with-routing-on-aws-s3/)
- [Minikube](https://kubernetes.io/docs/tutorials/hello-minikube/)

### About the Developer

**Colin Schofield**   
e: colin_sch@yahoo.com  
p: 0448 644 233  
l: https://www.linkedin.com/in/colins/
