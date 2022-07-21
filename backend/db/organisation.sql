
-- Table: Client
CREATE TABLE "Client" (
    Id serial  NOT NULL,
    FirstName varchar(20)  NOT NULL,
    LastName varchar(30)  NOT NULL,
    Patronymic varchar(30)  NOT NULL,
    Login varchar(30)  NOT NULL,
    Password varchar(30)  NOT NULL,
    Email varchar(256)  NULL,
    Telephone varchar(15)  NULL,
    EmailNotifications boolean  default false,
    GroupId int  NOT NULL,
    CONSTRAINT Clientpk PRIMARY KEY (Id)
);

-- Table: Test
CREATE TABLE "Test" (
    Id serial  NOT NULL,
    Title varchar(256)  NOT NULL,
    Text text  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NOT NULL,
    CONSTRAINT Testpk PRIMARY KEY (Id)
);

-- Table: Admin
CREATE TABLE "Admin" (
    Id serial  NOT NULL,
    FirstName varchar(20)  NOT NULL,
    LastName varchar(30)  NOT NULL,
    Login varchar(30)  NOT NULL,
    Password varchar(30)  NOT NULL,
    Email varchar(256)  NULL,
    CONSTRAINT Adminpk PRIMARY KEY (Id)
);

-- Table: Question
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

-- Table: Questionresult
CREATE TABLE "QuestionResult" (
	Id serial  NOT NULL,
	QuestionId int NOT null,
	ClientId int NOT null,
	Scores int not null,
	CONSTRAINT Questionresultpk PRIMARY KEY (Id)
);

-- Table: QuestionType
CREATE TABLE "QuestionType" (
    Id serial  NOT NULL,
    Type int  NOT NULL,
    CONSTRAINT QuestionTypepk PRIMARY KEY (Id)
);

-- Table: Answer
CREATE TABLE "Answer" (
    Id serial  NOT NULL,
    Text text  NOT NULL,
    Correct boolean  default false,
    QuestionId int8  NOT NULL,
    CONSTRAINT Answerpk PRIMARY KEY (Id)
);

-- Table: Course
CREATE TABLE "Course" (
    Id serial  NOT NULL,
    Title varchar(30)  NOT NULL,
    Date timestamp  NOT NULL,
    DateDel timestamp  NULL,
    ClientId int  NOT NULL,
    CONSTRAINT Coursepk PRIMARY KEY (Id)
);

--Table: Publicinfo
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


-- Table: Courseresults
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


-- Table: Department
CREATE TABLE "Department" (
    Id serial  NOT NULL,
    Title varchar(50)  NOT NULL,
    CONSTRAINT Departmentpk PRIMARY KEY (Id)
);

-- Table: Group
CREATE TABLE "Group" (
    Id serial  NOT NULL,
    Title varchar(20)  NOT NULL,
    DepartmentId int  NOT NULL,
    CONSTRAINT Grouppk PRIMARY KEY (Id)
);

-- Table: File
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

--Table: ActiveTest
CREATE TABLE "ActiveTest" (
    Id serial  NOT NULL,
    Start timestamp  NOT NULL,
    End timestamp  NOT NULL,
    Time timestamp  NOT NULL,
    Attempts smallint  NOT NULL,
    TestId int  NOT NULL,
    ClientId int  NOT NULL,
    Trainingtest boolean default false,
    CONSTRAINT ActiveTestpk PRIMARY KEY (Id)
);

-- Table: TestResults
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


-- foreign keys
-- Reference: QuestionresultClient (table: Questionresult)
ALTER TABLE "QuestionResult" ADD CONSTRAINT QuestionresultClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: PublicinfoClient (table: Publicinfo)
alter table "PublicInfo" add constraint PublicinfoClient
	FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: QuestionresultQuestion (table: Questionresult)
ALTER TABLE "QuestionResult" ADD CONSTRAINT QuestionresultQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;


-- Reference: ActiveTestClient (table: ActiveTest)
ALTER TABLE "ActiveTest" ADD CONSTRAINT ActiveTestClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: ActiveTestTest (table: Test)
ALTER TABLE "ActiveTest" ADD CONSTRAINT ActiveTestTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: AnswerQuestion (table: Answer)
ALTER TABLE "Answer" ADD CONSTRAINT AnswerQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: ClientGroup (table: Client)
ALTER TABLE "Client" ADD CONSTRAINT ClientGroup
    FOREIGN KEY (GroupId)
    REFERENCES "Group" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: CourseresultsClient (table: Courseresults)
ALTER TABLE "CourseResults" ADD CONSTRAINT CourseresultsClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: CourseresultsCreatedCourses (table: Courseresults)
ALTER TABLE "CourseResults" ADD CONSTRAINT CourseresultsCreatedCourses
    FOREIGN KEY (CourseId)
    REFERENCES "Course" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: CreatedCoursesClient (table: Course)
ALTER TABLE "Course" ADD CONSTRAINT CreatedCoursesClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;


-- Reference: FileClient (table: File)
ALTER TABLE "File" ADD CONSTRAINT FileClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: FileQuestion (table: File)
ALTER TABLE "File" ADD CONSTRAINT FileQuestion
    FOREIGN KEY (QuestionId)
    REFERENCES "Question" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: FileTest (table: File)
ALTER TABLE "File" ADD CONSTRAINT FileTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: GroupDepartment (table: Group)
ALTER TABLE "Group" ADD CONSTRAINT GroupDepartment
    FOREIGN KEY (DepartmentId)
    REFERENCES "Department" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: QuestionQuestionType (table: Question)
ALTER TABLE "Question" ADD CONSTRAINT QuestionQuestionType
    FOREIGN KEY (QuestionTypeId)
    REFERENCES "QuestionType" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: QuestionTest (table: Question)
ALTER TABLE "Question" ADD CONSTRAINT QuestionTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: TestResultsClient (table: TestResults)
ALTER TABLE "TestResults" ADD CONSTRAINT TestResultsClient
    FOREIGN KEY (ClientId)
    REFERENCES "Client" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: TestResultsCourse (table: TestResults)
ALTER TABLE "TestResults" ADD CONSTRAINT TestResultsCourse
    FOREIGN KEY (CourseId)
    REFERENCES "Course" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: TestResultsTest (table: TestResults)
ALTER TABLE "TestResults" ADD CONSTRAINT TestResultsTest
    FOREIGN KEY (TestId)
    REFERENCES "Test" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;


