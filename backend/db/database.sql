-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2022-07-07 13:20:23.742

-- tables
-- Table: Active_Test
CREATE TABLE Active_Test (
    Id bigserial  NOT NULL,
    Start timestamp  NOT NULL,
    "End" timestamp  NOT NULL,
    Time timestamp  NOT NULL,
    Attempts smallint  NOT NULL,
    Access_Key varchar(20)  NOT NULL,
    Client_Id bigint  NOT NULL,
    CONSTRAINT Active_Test_pk PRIMARY KEY (Id)
);

-- Table: Admin
CREATE TABLE Admin (
    Id bigserial  NOT NULL,
    First_Name varchar(20)  NOT NULL,
    Last_Name varchar(30)  NOT NULL,
    Login varchar(30)  NOT NULL,
    Password varchar(30)  NOT NULL,
    Email varchar(256)  NULL,
    CONSTRAINT Admin_pk PRIMARY KEY (Id)
);

-- Table: Answer
CREATE TABLE Answer (
    Id bigserial  NOT NULL,
    Text text  NOT NULL,
    Correct boolean  NOT NULL,
    Question_Id int8  NOT NULL,
    CONSTRAINT Answer_pk PRIMARY KEY (Id)
);

-- Table: Client
CREATE TABLE Client (
    Id bigserial  NOT NULL,
    First_Name varchar(20)  NOT NULL,
    Last_Name varchar(30)  NOT NULL,
    Patronymic varchar(30)  NOT NULL,
    Login varchar(30)  NOT NULL,
    Password varchar(30)  NOT NULL,
    Email varchar(256)  NULL,
    Telephone varchar(15)  NULL,
    Email_Notifications boolean  NOT NULL,
    Level_Id smallint  NOT NULL,
    Group_Id int  NOT NULL,
    Organization_Id int8  NOT NULL,
    CONSTRAINT Client_pk PRIMARY KEY (Id)
);

-- Table: Course
CREATE TABLE Course (
    Id bigserial  NOT NULL,
    Tittle varchar(30)  NOT NULL,
    Date timestamp  NOT NULL,
    Date_Del timestamp  NULL,
    Client_Id bigint  NOT NULL,
    Type_Entity_Id int  NOT NULL,
    CONSTRAINT Course_pk PRIMARY KEY (Id)
);

-- Table: Course_results
CREATE TABLE Course_results (
    Id serial  NOT NULL,
    Time time  NOT NULL,
    Date date  NOT NULL,
    Assessment varchar(20)  NOT NULL,
    Scores real  NOT NULL,
    Client_Id bigint  NOT NULL,
    Course_Id bigint  NOT NULL,
    CONSTRAINT Course_results_pk PRIMARY KEY (Id)
);

-- Table: Dead_Line
CREATE TABLE Dead_Line (
    Id serial  NOT NULL,
    Date timestamp  NOT NULL,
    Level_Id smallint  NOT NULL,
    Organization_Id int8  NOT NULL,
    CONSTRAINT Dead_Line_pk PRIMARY KEY (Id)
);

-- Table: Department
CREATE TABLE Department (
    Id serial  NOT NULL,
    Tittle varchar(50)  NOT NULL,
    Organization_Id int8  NOT NULL,
    CONSTRAINT Department_pk PRIMARY KEY (Id)
);

-- Table: File
CREATE TABLE File (
    Id bigserial  NOT NULL,
    Date timestamp  NOT NULL,
    Date_Del timestamp  NULL,
    File_Name varchar(256)  NOT NULL,
    File_Name_Tmp varchar(256)  NOT NULL,
    Test_Id int8  NOT NULL,
    Question_Id bigint  NOT NULL,
    Client_Id int8  NOT NULL,
    CONSTRAINT File_pk PRIMARY KEY (Id)
);

-- Table: Group
CREATE TABLE "Group" (
    Id serial  NOT NULL,
    Title varchar(20)  NOT NULL,
    Department_Id int  NOT NULL,
    CONSTRAINT Group_pk PRIMARY KEY (Id)
);

-- Table: Level
CREATE TABLE "Level" (
    Id smallint  NOT NULL,
    Title varchar(100)  NOT NULL,
    Price int  NOT NULL,
    Paid boolean  NOT NULL,
    Create_Course boolean  NOT NULL,
    Take_Course boolean  NOT NULL,
    Apload_File boolean  NOT NULL,
    View_Your_Result boolean  NOT NULL,
    View_Other_Results boolean  NOT NULL,
    CONSTRAINT Level_pk PRIMARY KEY (Id)
);

-- Table: Organization
CREATE TABLE Organization (
    Id bigserial  NOT NULL,
    Title int  NOT NULL,
    Login varchar(30)  NOT NULL,
    Password varchar(30)  NOT NULL,
    Email varchar(256)  NULL,
    Telephone varchar(15)  NULL,
    Email_Notifications boolean  NOT NULL,
    Level_Id smallint  NOT NULL,
    CONSTRAINT Organization_pk PRIMARY KEY (Id)
);

-- Table: Payment
CREATE TABLE Payment (
    Number int  NOT NULL,
    Name varchar(100)  NOT NULL,
    Money int  NOT NULL,
    Date timestamp  NOT NULL,
    Level_Id smallint  NOT NULL,
    Client_Id int8  NOT NULL,
    CONSTRAINT Payment_pk PRIMARY KEY (Number)
);

-- Table: Question
CREATE TABLE Question (
    Id bigserial  NOT NULL,
    Text text  NOT NULL,
    Date timestamp  NOT NULL,
    Date_Del timestamp  NOT NULL,
    Type_Entity_Id int  NOT NULL,
    Test_Id bigint  NOT NULL,
    Question_Type_Id int  NOT NULL,
    CONSTRAINT Question_pk PRIMARY KEY (Id)
);

-- Table: Question_Type
CREATE TABLE Question_Type (
    Id serial  NOT NULL,
    Type int  NOT NULL,
    CONSTRAINT Question_Type_pk PRIMARY KEY (Id)
);

-- Table: Session
CREATE TABLE Session (
    Id int  NOT NULL,
    Date int  NOT NULL,
    Token int  NOT NULL,
    Client_Id int8  NOT NULL,
    CONSTRAINT Session_pk PRIMARY KEY (Id)
);

-- Table: Test
CREATE TABLE Test (
    Id bigserial  NOT NULL,
    Title varchar(256)  NOT NULL,
    Text text  NOT NULL,
    Date timestamp  NOT NULL,
    Date_Del timestamp  NOT NULL,
    Type_Entity_Id int  NOT NULL,
    CONSTRAINT Test_pk PRIMARY KEY (Id)
);

-- Table: Test_Results
CREATE TABLE Test_Results (
    Id bigserial  NOT NULL,
    Time time  NOT NULL,
    Date date  NOT NULL,
    Client_Id bigint  NOT NULL,
    Test_Id int8  NOT NULL,
    Assessment varchar(20)  NOT NULL,
    Passage_Time time  NOT NULL,
    Scores real  NOT NULL,
    Course_Id int8  NOT NULL,
    CONSTRAINT Test_Results_pk PRIMARY KEY (Id)
);

-- Table: Trash
CREATE TABLE Trash (
    Id bigserial  NOT NULL,
    Client_Id bigint  NOT NULL,
    Record_ID bigint  NOT NULL,
    Type_Entity_Id int  NOT NULL,
    Date_Del timestamp  NOT NULL,
    CONSTRAINT Trash_pk PRIMARY KEY (Id)
);

-- Table: Type_Entity
CREATE TABLE Type_Entity (
    Id serial  NOT NULL,
    Type varchar(256)  NOT NULL,
    CONSTRAINT Type_Entity_pk PRIMARY KEY (Id)
);

-- foreign keys
-- Reference: Active_Test_Client (table: Active_Test)
ALTER TABLE Active_Test ADD CONSTRAINT Active_Test_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Answer_Question (table: Answer)
ALTER TABLE Answer ADD CONSTRAINT Answer_Question
    FOREIGN KEY (Question_Id)
    REFERENCES Question (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Client_Group (table: Client)
ALTER TABLE Client ADD CONSTRAINT Client_Group
    FOREIGN KEY (Group_Id)
    REFERENCES "Group" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Client_Level (table: Client)
ALTER TABLE Client ADD CONSTRAINT Client_Level
    FOREIGN KEY (Level_Id)
    REFERENCES "Level" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Client_Organization (table: Client)
ALTER TABLE Client ADD CONSTRAINT Client_Organization
    FOREIGN KEY (Organization_Id)
    REFERENCES Organization (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Course_results_Client (table: Course_results)
ALTER TABLE Course_results ADD CONSTRAINT Course_results_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Course_results_Created_Courses (table: Course_results)
ALTER TABLE Course_results ADD CONSTRAINT Course_results_Created_Courses
    FOREIGN KEY (Course_Id)
    REFERENCES Course (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Created_Courses_Client (table: Course)
ALTER TABLE Course ADD CONSTRAINT Created_Courses_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Created_Courses_Type_Entity (table: Course)
ALTER TABLE Course ADD CONSTRAINT Created_Courses_Type_Entity
    FOREIGN KEY (Type_Entity_Id)
    REFERENCES Type_Entity (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Dead_Line_Organization (table: Dead_Line)
ALTER TABLE Dead_Line ADD CONSTRAINT Dead_Line_Organization
    FOREIGN KEY (Organization_Id)
    REFERENCES Organization (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Department_Organization (table: Department)
ALTER TABLE Department ADD CONSTRAINT Department_Organization
    FOREIGN KEY (Organization_Id)
    REFERENCES Organization (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: File_Client (table: File)
ALTER TABLE File ADD CONSTRAINT File_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: File_Question (table: File)
ALTER TABLE File ADD CONSTRAINT File_Question
    FOREIGN KEY (Question_Id)
    REFERENCES Question (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: File_Test (table: File)
ALTER TABLE File ADD CONSTRAINT File_Test
    FOREIGN KEY (Test_Id)
    REFERENCES Test (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Group_Department (table: Group)
ALTER TABLE "Group" ADD CONSTRAINT Group_Department
    FOREIGN KEY (Department_Id)
    REFERENCES Department (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Organisation_Organisation_Level (table: Dead_Line)
ALTER TABLE Dead_Line ADD CONSTRAINT Organisation_Organisation_Level
    FOREIGN KEY (Level_Id)
    REFERENCES "Level" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Payment_Client (table: Payment)
ALTER TABLE Payment ADD CONSTRAINT Payment_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Payment_Level (table: Payment)
ALTER TABLE Payment ADD CONSTRAINT Payment_Level
    FOREIGN KEY (Level_Id)
    REFERENCES "Level" (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Question_Question_Type (table: Question)
ALTER TABLE Question ADD CONSTRAINT Question_Question_Type
    FOREIGN KEY (Question_Type_Id)
    REFERENCES Question_Type (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Question_Test (table: Question)
ALTER TABLE Question ADD CONSTRAINT Question_Test
    FOREIGN KEY (Test_Id)
    REFERENCES Test (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Question_Type_Entity (table: Question)
ALTER TABLE Question ADD CONSTRAINT Question_Type_Entity
    FOREIGN KEY (Type_Entity_Id)
    REFERENCES Type_Entity (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Session_Client (table: Session)
ALTER TABLE Session ADD CONSTRAINT Session_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Test_Results_Client (table: Test_Results)
ALTER TABLE Test_Results ADD CONSTRAINT Test_Results_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Test_Results_Course (table: Test_Results)
ALTER TABLE Test_Results ADD CONSTRAINT Test_Results_Course
    FOREIGN KEY (Course_Id)
    REFERENCES Course (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Test_Results_Test (table: Test_Results)
ALTER TABLE Test_Results ADD CONSTRAINT Test_Results_Test
    FOREIGN KEY (Test_Id)
    REFERENCES Test (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Test_Type_Entity (table: Test)
ALTER TABLE Test ADD CONSTRAINT Test_Type_Entity
    FOREIGN KEY (Type_Entity_Id)
    REFERENCES Type_Entity (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Trash_Client (table: Trash)
ALTER TABLE Trash ADD CONSTRAINT Trash_Client
    FOREIGN KEY (Client_Id)
    REFERENCES Client (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Trash_Type_Entity (table: Trash)
ALTER TABLE Trash ADD CONSTRAINT Trash_Type_Entity
    FOREIGN KEY (Type_Entity_Id)
    REFERENCES Type_Entity (Id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.

