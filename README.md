# GO REST 
This project will only use std-lib to implement all requirements

## Motivation
The purpose of this project is to practise Golang using only std-lib with eclipse/Symphony code structure 
and architecture. 
This project is **NOT** meant for learning Golang basics - syntax etc. the idea is to create project
that allow you to get fluent using Golang having close to real world problems to solve. 
Simply take below requirements and get familiar with eclipse/Symphony architecture 
(https://www.linkedin.com/pulse/hb-mvp-design-pattern-extensible-systems-part-i-haishi-bai/) and start solving problem
by **YOUR OWN** take a look at my solution only when necessary. At /design path are located sequence diagrams to 
illustrate basic flows like bootstrapping or get resource by id. 

## Requirements
### Overview
Main purpose of this project will be to allow manage Components that are part of eclipse/Symphony Solution.
Symphony primarily does not support Components separately only as a part of Solution. 
### Architecture and Implementation 
Whole application have to be implemented according to HB-MVP design pattern the same way as Symphony except that code base will not be the same as it's only for learning purposes. 
For implementation only standard library should be used no frameworks, lib etc. 
For persistence layer PostgresSQL should be implemented. 
### Features
- HTTP binding in REST style
- CRUD for Components
- At start create DB structure
- PostgresSQL as underlying DB
- Test for each layer 
- Logging to file and possibility to set other output

