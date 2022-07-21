CREATE TABLE "Client" (
    Id serial  NOT NULL,
    FirstName varchar(20)  NOT NULL,
    LastName varchar(30)  NOT NULL,
    Patronymic varchar(30)  NOT NULL,
    Password varchar(60)  NOT NULL,
    Email varchar(256)  NOT NULL,
    Telephone varchar(15)  NULL,
    EmailNotifications boolean  default false,
    GroupId int  NULL,
    CONSTRAINT Clientpk PRIMARY KEY (Id)
);

CREATE TABLE "Test" (
    Id serial  NOT NULL,
    Title varchar(256)  NOT NULL,
    Text text  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NULL,
    CONSTRAINT Testpk PRIMARY KEY (Id)
);

CREATE TABLE "Admin" (
    Id serial  NOT NULL,
    FirstName varchar(20)  NOT NULL,
    LastName varchar(30)  NOT NULL,
    Password varchar(60)  NOT NULL,
    Email varchar(256)  NULL,
    CONSTRAINT Adminpk PRIMARY KEY (Id)
);

CREATE TABLE "Question" (
    Id serial  NOT NULL,
    Text text  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NOT NULL,
    TestId int  NOT NULL,
    QuestionTypeId int  NOT NULL,
    Hint varchar(150),
    CONSTRAINT Questionpk PRIMARY KEY (Id)
);

CREATE TABLE "QuestionResult" (
	Id serial  NOT NULL,
	QuestionId int NOT null,
	ClientId int NOT null,
	Scores int NOT null,
	CONSTRAINT Questionresultpk PRIMARY KEY (Id)
);

CREATE TABLE "QuestionType" (
    Id serial  NOT NULL,
    Type int  NOT NULL,
    CONSTRAINT QuestionTypepk PRIMARY KEY (Id)
);

CREATE TABLE "Answer" (
    Id serial  NOT NULL,
    Text text  NOT NULL,
    Correct boolean  default false,
    QuestionId int8  NOT NULL,
    CONSTRAINT Answerpk PRIMARY KEY (Id)
);

CREATE TABLE "Course" (
    Id serial  NOT NULL,
    Title varchar(30)  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NULL,
    ClientId int  NOT NULL,
    CONSTRAINT Coursepk PRIMARY KEY (Id)
);

CREATE TABLE "PublicInfo" (
    Id serial  NOT NULL,
    Title varchar(30)  NOT NULL,
    Text text,
    File varchar(200),
    Date timestamp  NULL,
    DateDel timestamp  NULL,
    ClientId int  NOT NULL,
    CONSTRAINT Publicinfopk PRIMARY KEY (Id)
);

CREATE TABLE "CourseResults" (
    Id serial  NOT NULL,
    Time time  NOT NULL,
    Date date  NOT NULL,
    Assessment varchar(20)  NOT NULL,
    Scores real  NOT NULL,
    ClientId int  NOT NULL,
    CourseId int  NOT NULL,
    CONSTRAINT Courseresultspk PRIMARY KEY (Id)
);

CREATE TABLE "Department" (
    Id serial  NOT NULL,
    Title varchar(50)  NOT NULL,
    CONSTRAINT Departmentpk PRIMARY KEY (Id)
);

CREATE TABLE "Group" (
    Id serial  NOT NULL,
    Title varchar(20)  NOT NULL,
    DepartmentId int  NOT NULL,
    CONSTRAINT Grouppk PRIMARY KEY (Id)
);

CREATE TABLE "File" (
    Id serial  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NULL,
    FileName varchar(256)  NOT NULL,
    FileNameTmp varchar(256)  NOT NULL,
    TestId int8  NOT NULL,
    QuestionId int  NOT NULL,
    ClientId int8  NOT NULL,
    CONSTRAINT Filepk PRIMARY KEY (Id)
);

CREATE TABLE "ActiveTest" (
    Id serial  NOT NULL,
    TimeStart timestamp  NOT NULL,
    TimeEnd timestamp  NOT NULL,
    Time timestamp  NOT NULL,
    Attempts smallint  NOT NULL,
    TestId int  NOT NULL,
    ClientId int  NOT NULL,
    Trainingtest boolean default false,
    CONSTRAINT ActiveTestpk PRIMARY KEY (Id)
);

CREATE TABLE "TestResults" (
    Id serial  NOT NULL,
    Time time  NOT NULL,
    Date date  NOT NULL,
    ClientId int  NOT NULL,
    TestId int8  NOT NULL,
    Assessment varchar(20)  NOT NULL,
    PassageTime time  NOT NULL,
    Scores real  NOT NULL,
    CourseId int8  NOT NULL,
    CONSTRAINT TestResultspk PRIMARY KEY (Id)
);

ALTER TABLE "QuestionResult" ADD CONSTRAINT QuestionresultClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

alter table "PublicInfo" add constraint PublicinfoClient
	FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "QuestionResult" ADD CONSTRAINT QuestionresultQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "ActiveTest" ADD CONSTRAINT ActiveTestClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "ActiveTest" ADD CONSTRAINT ActiveTestTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Answer" ADD CONSTRAINT AnswerQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Client" ADD CONSTRAINT ClientGroup
    FOREIGN KEY (GroupId)
    REFERENCES "Group" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "CourseResults" ADD CONSTRAINT CourseresultsClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "CourseResults" ADD CONSTRAINT CourseresultsCreatedCourses
    FOREIGN KEY (CourseId)
    REFERENCES "Course" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Course" ADD CONSTRAINT CreatedCoursesClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "File" ADD CONSTRAINT FileClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "File" ADD CONSTRAINT FileQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "File" ADD CONSTRAINT FileTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Group" ADD CONSTRAINT GroupDepartment
    FOREIGN KEY (DepartmentId)
    REFERENCES "Department" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Question" ADD CONSTRAINT QuestionQuestionType
    FOREIGN KEY (QuestionTypeId)
    REFERENCES "QuestionType" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Question" ADD CONSTRAINT QuestionTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "TestResults" ADD CONSTRAINT TestResultsClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "TestResults" ADD CONSTRAINT TestResultsCourse
    FOREIGN KEY (CourseId)
    REFERENCES "Course" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "TestResults" ADD CONSTRAINT TestResultsTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;


