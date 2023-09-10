# Seller Template Project
Seller Project, Developed using GoLang

In this sample project, I have developed a part of a web store api as a template through Golang language.</br>
To develop this project, I have used clean architecture, which includes the following main layers.</br>
<ul>
<li>Domain Layer: Entities and Services Model</li>
<li>Application: Project Algorithms and Process</li>
<li>Infrastructure: Implemention of Interfaces that Presented in Application Layer</li>
<li>Presentation (final web service): Api, Controllers and DI Handlers</li>
</ul>
The image below shows the layers and their dependencies.
<br>
<img style="width:500px;" src="https://github.com/rabieedev1996/Seller/assets/20272337/7bae4472-f900-4ed3-9a89-aa121fbf574e"></img>
<br>
<a src="https://github.com/rabieedev1996/Seller/assets/20272337/e4ea8231-26d7-4491-8c40-7a16f83f2f9a">View Full Size</a>

<h1>Request Flow</h1>
The flow of executing a request is generally as follows from start to finish.
<br>
<img  src="https://github.com/rabieedev1996/Seller/assets/20272337/c679e44b-11ba-4380-8300-74be734d6989"></img>
<br>
<a src="https://github.com/rabieedev1996/Seller/assets/20272337/5b9275ee-421c-462d-9b6a-e6e67b338ae1">View Full Size</a>

<h1>Dependency Inversion and Dependency Injection</h1>

In general, in order to separate the implementation of external elements such as the database or SMS service from the business, it is good to use the Dependency Inversion. In this way, the application layer uses an interface injected by the API layer. This interface is implemented in the infrastructure layer. The task of creating the correct object and sending it to the application is the responsibility of the Api layer, which is called dependency injection. 
You can see the diagram of this part in the picture below.

<img  src="https://github.com/rabieedev1996/Seller/assets/20272337/0f50e030-9439-4dc2-a2f5-bd6040c50639"></img>
<a src="https://github.com/rabieedev1996/Seller/assets/20272337/9bd4e50e-9dcf-424a-8d00-f66ee809315b">View Full Size</a>

Finally, I mention all the significant tools and frameworks used in this project in the form of keywords.

<ul>
<li>Clean Architecture</li>
<li>GIN Framework</li>
<li>GORM ORM</li>
<li>Generic repositoy pattern</li>
<li>CQRS pattern</li>
<li>PostgreSql DB</li>
<li>MongoDB</li>
<li>Redis</li>
<li>Swagger Documentation</li>
<li>Dependency Injection</li>
<li>Dependency Inversion</li>

</ul>
