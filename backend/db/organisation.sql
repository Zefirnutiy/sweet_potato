
-- Table: Client
CREATE TABLE "KTK"."Client" (
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

-- Table: Test
CREATE TABLE "KTK"."Test" (
    Id serial  NOT NULL,
    Title varchar(256)  NOT NULL,
    Text text  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NULL,
    CONSTRAINT Testpk PRIMARY KEY (Id)
);

-- Table: Admin
CREATE TABLE "KTK"."Admin" (
    Id serial  NOT NULL,
    FirstName varchar(20)  NOT NULL,
    LastName varchar(30)  NOT NULL,
    Password varchar(60)  NOT NULL,
    Email varchar(256)  NULL,
    CONSTRAINT Adminpk PRIMARY KEY (Id)
);

-- Table: Question
CREATE TABLE "KTK"."Question" (
    Id serial  NOT NULL,
    Text text  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NOT NULL,
    TestId int  NOT NULL,
    QuestionTypeId int  NOT NULL,
    Hint varchar(150),
    CONSTRAINT Questionpk PRIMARY KEY (Id)
);

-- Table: Questionresult
CREATE TABLE "KTK"."Questionresult" (
	Id serial  NOT NULL,
	QuestionId int NOT null,
	ClientId int NOT null,
	Scores int NOT null,
	CONSTRAINT Questionresultpk PRIMARY KEY (Id)
);

-- Table: QuestionType
CREATE TABLE "KTK"."QuestionType" (
    Id serial  NOT NULL,
    Type int  NOT NULL,
    CONSTRAINT QuestionTypepk PRIMARY KEY (Id)
);

-- Table: Answer
CREATE TABLE "KTK"."Answer" (
    Id serial  NOT NULL,
    Text text  NOT NULL,
    Correct boolean  default false,
    QuestionId int8  NOT NULL,
    CONSTRAINT Answerpk PRIMARY KEY (Id)
);

-- Table: Course
CREATE TABLE "KTK"."Course" (
    Id serial  NOT NULL,
    Title varchar(30)  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NULL,
    ClientId int  NOT NULL,
    CONSTRAINT Coursepk PRIMARY KEY (Id)
);

--Table: Publicinfo
CREATE TABLE "KTK"."Publicinfo" (
    Id serial  NOT NULL,
    Title varchar(30)  NOT NULL,
    Text text,
    File varchar(200),
    Date timestamp  NULL,
    DateDel timestamp  NULL,
    ClientId int  NOT NULL,
    CONSTRAINT Publicinfopk PRIMARY KEY (Id)
);


-- Table: Courseresults
CREATE TABLE "KTK"."Courseresults" (
    Id serial  NOT NULL,
    Time time  NOT NULL,
    Date date  NOT NULL,
    Assessment varchar(20)  NOT NULL,
    Scores real  NOT NULL,
    ClientId int  NOT NULL,
    CourseId int  NOT NULL,
    CONSTRAINT Courseresultspk PRIMARY KEY (Id)
);


-- Table: Department
CREATE TABLE "KTK"."Department" (
    Id serial  NOT NULL,
    Title varchar(50)  NOT NULL,
    CONSTRAINT Departmentpk PRIMARY KEY (Id)
);

-- Table: Group
CREATE TABLE "KTK"."Group" (
    Id serial  NOT NULL,
    Title varchar(20)  NOT NULL,
    DepartmentId int  NOT NULL,
    CONSTRAINT Grouppk PRIMARY KEY (Id)
);

-- Table: File
CREATE TABLE "KTK"."File" (
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

--Table: ActiveTest
CREATE TABLE "KTK"."ActiveTest" (
    Id serial  NOT NULL,
    Start timestamp  NOT NULL,
    "End" timestamp  NOT NULL,
    Time timestamp  NOT NULL,
    Attempts smallint  NOT NULL,
    TestId int  NOT NULL,
    ClientId int  NOT NULL,
    Trainingtest boolean default false,
    CONSTRAINT ActiveTestpk PRIMARY KEY (Id)
);

-- Table: TestResults
CREATE TABLE "KTK"."TestResults" (
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


-- foreign keys
-- Reference: QuestionresultClient (table: Questionresult)
ALTER TABLE "KTK"."Questionresult" ADD CONSTRAINT QuestionresultClient
    FOREIGN KEY (ClientId)
    REFERENCES "KTK"."Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: PublicinfoClient (table: Publicinfo)
alter table "KTK"."Publicinfo" add constraint PublicinfoClient
	FOREIGN KEY (ClientId)
    REFERENCES "KTK"."Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: QuestionresultQuestion (table: Questionresult)
ALTER TABLE "KTK"."Questionresult" ADD CONSTRAINT QuestionresultQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "KTK"."Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;


-- Reference: ActiveTestClient (table: ActiveTest)
ALTER TABLE "KTK"."ActiveTest" ADD CONSTRAINT ActiveTestClient
    FOREIGN KEY (ClientId)
    REFERENCES "KTK"."Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: ActiveTestTest (table: Test)
ALTER TABLE "KTK"."ActiveTest" ADD CONSTRAINT ActiveTestTest
    FOREIGN KEY (TestId)
    REFERENCES "KTK"."Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: AnswerQuestion (table: Answer)
ALTER TABLE "KTK"."Answer" ADD CONSTRAINT AnswerQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "KTK"."Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: ClientGroup (table: Client)
ALTER TABLE "KTK"."Client" ADD CONSTRAINT ClientGroup
    FOREIGN KEY (GroupId)
    REFERENCES "KTK"."Group" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: CourseresultsClient (table: Courseresults)
ALTER TABLE "KTK"."Courseresults" ADD CONSTRAINT CourseresultsClient
    FOREIGN KEY (ClientId)
    REFERENCES "KTK"."Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: CourseresultsCreatedCourses (table: Courseresults)
ALTER TABLE "KTK"."Courseresults" ADD CONSTRAINT CourseresultsCreatedCourses
    FOREIGN KEY (CourseId)
    REFERENCES "KTK"."Course" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: CreatedCoursesClient (table: Course)
ALTER TABLE "KTK"."Course" ADD CONSTRAINT CreatedCoursesClient
    FOREIGN KEY (ClientId)
    REFERENCES "KTK"."Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;


-- Reference: FileClient (table: File)
ALTER TABLE "KTK"."File" ADD CONSTRAINT FileClient
    FOREIGN KEY (ClientId)
    REFERENCES "KTK"."Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: FileQuestion (table: File)
ALTER TABLE "KTK"."File" ADD CONSTRAINT FileQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "KTK"."Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: FileTest (table: File)
ALTER TABLE "KTK"."File" ADD CONSTRAINT FileTest
    FOREIGN KEY (TestId)
    REFERENCES "KTK"."Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: GroupDepartment (table: Group)
ALTER TABLE "KTK"."Group" ADD CONSTRAINT GroupDepartment
    FOREIGN KEY (DepartmentId)
    REFERENCES "KTK"."Department" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: QuestionQuestionType (table: Question)
ALTER TABLE "KTK"."Question" ADD CONSTRAINT QuestionQuestionType
    FOREIGN KEY (QuestionTypeId)
    REFERENCES "KTK"."QuestionType" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: QuestionTest (table: Question)
ALTER TABLE "KTK"."Question" ADD CONSTRAINT QuestionTest
    FOREIGN KEY (TestId)
    REFERENCES "KTK"."Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: TestResultsClient (table: TestResults)
ALTER TABLE "KTK"."TestResults" ADD CONSTRAINT TestResultsClient
    FOREIGN KEY (ClientId)
    REFERENCES "KTK"."Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: TestResultsCourse (table: TestResults)
ALTER TABLE "KTK"."TestResults" ADD CONSTRAINT TestResultsCourse
    FOREIGN KEY (CourseId)
    REFERENCES "KTK"."Course" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: TestResultsTest (table: TestResults)
ALTER TABLE "KTK"."TestResults" ADD CONSTRAINT TestResultsTest
    FOREIGN KEY (TestId)
    REFERENCES "KTK"."Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;


