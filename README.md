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
In the image below, an overview of the architecture and components of the project is specified.</br>
Postgresql is also used as  database. To connect to the database, I have used the repository pattern with generic, which is based on the implementation of GORM ORM and Raw Query.</br>
Also, the Gin framework has been used for routing in the project

<img style="margin-top:20px;" src="https://classicode.org/seller.png" />
