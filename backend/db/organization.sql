-- CREATE SCHEMA <>;
-- set schema '<>';
CREATE TABLE IF NOT EXISTS "Client"
(
    "Id" serial NOT NULL,
    "FirstName" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "LastName" character varying(30) COLLATE pg_catalog."default" NOT NULL,
    "Patronymic" character varying(30) COLLATE pg_catalog."default" NOT NULL,
    "Password" character varying(256) COLLATE pg_catalog."default" NOT NULL,
    "Email" character varying(256) COLLATE pg_catalog."default" NOT NULL,
    "Telephone" character varying(15) COLLATE pg_catalog."default",
    "EmailNotifications" boolean DEFAULT false,
    "ManageCourses" boolean DEFAULT false,
    "ManageUsers" boolean DEFAULT false,
    "AploadFiles" boolean DEFAULT false,
    "ViewYourResults" boolean DEFAULT false,
    "ViewOtherResults" boolean DEFAULT false,
    "DepartmentId" integer NOT NULL,
    "GroupId" integer NOT NULL,
    "CreatorId" integer,
    CONSTRAINT clientpk PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "Course"
(
    "Id" serial NOT NULL,
    "Title" character varying(30) COLLATE pg_catalog."default" NOT NULL,
    "Text" text COLLATE pg_catalog."default" NOT NULL,
    "Files" boolean NOT NULL DEFAULT false,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "DateDel" date,
    "TimeDel" time without time zone,
    "ClientId" integer NOT NULL,
    CONSTRAINT "CoursePk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "CourseResults"
(
    "Id" serial NOT NULL,
    "Assessment" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Scores" real NOT NULL,
    "NumberTests" integer NOT NULL DEFAULT 0,
    "PassageTime" time without time zone NOT NULL,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "ClientId" integer NOT NULL,
    "CourseId" integer NOT NULL,
    CONSTRAINT "CourseResultsPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "ActiveCourse"
(
    "Id" serial NOT NULL,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "DateClose" date NOT NULL,
    "TimeClose" time without time zone NOT NULL,
    "CourseId" integer NOT NULL,
    "ClientId" integer NOT NULL,
    CONSTRAINT "ActiveCoursePk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "Department"
(
    "Id" serial NOT NULL,
    "Title" character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "DepartmentPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "Group"
(
    "Id" serial NOT NULL,
    "Title" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "TitleSingular" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "DepartmentId" integer NOT NULL,
    CONSTRAINT "GroupPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "File"
(
    "Id" serial NOT NULL,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "DateDel" date,
    "TimeDel" time without time zone,
    "FileName" character varying(256) COLLATE pg_catalog."default" NOT NULL,
    "FileNameTmp" character varying(256) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "FilePk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "FileInformation"
(
    "ClientId" bigint NOT NULL,
    "FileId" bigint NOT NULL,
    "PublicInfoId" bigint,
    "TestId" bigint,
    "QuestionId" bigint,
    CONSTRAINT "FileInformationPk" PRIMARY KEY ("FileId")
);

CREATE TABLE IF NOT EXISTS "PublicInfo"
(
    "Id" serial NOT NULL,
    "Title" character varying(30) COLLATE pg_catalog."default" NOT NULL,
    "Text" text COLLATE pg_catalog."default",
    "Files" boolean NOT NULL DEFAULT false,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "DateDel" date,
    "TimeDel" time without time zone,
    "ClientId" integer NOT NULL,
    CONSTRAINT "PublicInfoPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "Test"
(
    "Id" serial NOT NULL,
    "Title" character varying(256) COLLATE pg_catalog."default" NOT NULL,
    "Text" text COLLATE pg_catalog."default" NOT NULL,
    "Files" boolean NOT NULL DEFAULT false,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "DateDel" date,
    "TimeDel" time without time zone,
    CONSTRAINT "TestPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "TestResults"
(
    "Id" serial NOT NULL,
    "PassageTime" time without time zone NOT NULL,
    "Assessment" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Scores" real NOT NULL,
    "Attempts" smallint NOT NULL,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "CourseId" bigint NOT NULL,
    "TestId" bigint NOT NULL,
    "ClientId" bigint NOT NULL,
    CONSTRAINT "TestResultsPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "ActiveTest"
(
    "Id" serial NOT NULL,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "DateClose" date NOT NULL,
    "TimeClose" time without time zone NOT NULL,
    "Attempts" smallint NOT NULL,
    "TestTypeId" smallint NOT NULL DEFAULT 1,
    "TestId" bigint NOT NULL,
    "ClientId" bigint NOT NULL,
    CONSTRAINT "ActiveTestPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "TestType"
(
    "Id" serial NOT NULL,
    "Title" integer NOT NULL,
    CONSTRAINT "TestTypePk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "Question"
(
    "Id" serial NOT NULL,
    "Text" text COLLATE pg_catalog."default" NOT NULL,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "DateDel" date,
    "TimeDel" time without time zone,
    "Hint" character varying(150) COLLATE pg_catalog."default",
    "AnswerVariant" text NOT NULL, --Варианты ответа будут храниться в таком виде - ответ;ответ;ответ;ответ
    "AnswerCorrect" text NOT NULL, --Верные варианты ответа будут храниться в таком виде - ответ;ответ;ответ;ответ
    "Files" boolean NOT NULL DEFAULT false,
    "TestId" bigint NOT NULL,
    "QuestionTypeId" smallint NOT NULL,
    CONSTRAINT "QuestionPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "QuestionResult"
(
    "Id" serial NOT NULL,
    "Date" date NOT NULL,
    "Time" time without time zone NOT NULL,
    "Scores" integer NOT NULL,
    "QuestionId" integer NOT NULL,
    "ClientId" integer NOT NULL,
    CONSTRAINT "QuestionResultPk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "QuestionType"
(
    "Id" serial NOT NULL,
    "Title" integer NOT NULL,
    "InputTypeId" integer NOT NULL default 1,
    CONSTRAINT "QuestionTypePk" PRIMARY KEY ("Id")
);

CREATE TABLE IF NOT EXISTS "InputType"
(
    "Id" serial NOT NULL,
    "Title" character varying(50) NOT NULL,
    "Type" character varying(20) NOT NULL,
    CONSTRAINT "InputTypePk" PRIMARY KEY ("Id")
);

--////////////////////

ALTER TABLE IF EXISTS "Client"
    ADD CONSTRAINT "ClientCreator" FOREIGN KEY ("CreatorId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "Client"
    ADD CONSTRAINT "ClientDepartment" FOREIGN KEY ("DepartmentId")
    REFERENCES "Department" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "Client"
    ADD CONSTRAINT "ClientGroup" FOREIGN KEY ("GroupId")
    REFERENCES "Group" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "Course"
    ADD CONSTRAINT "CreatedCoursesClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "CourseResults"
    ADD CONSTRAINT "CourseResultsClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "CourseResults"
    ADD CONSTRAINT "CourseResultsCreatedCourses" FOREIGN KEY ("CourseId")
    REFERENCES "Course" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "ActiveCourse"
    ADD CONSTRAINT "ActiveCourseCourse" FOREIGN KEY ("CourseId")
    REFERENCES "Course" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "ActiveCourse"
    ADD CONSTRAINT "ActiveCourseClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "Group"
    ADD CONSTRAINT "GroupDepartment" FOREIGN KEY ("DepartmentId")
    REFERENCES "Department" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "FileInformation"
    ADD CONSTRAINT "FileInformationClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "FileInformation"
    ADD CONSTRAINT "FileInformationFile" FOREIGN KEY ("FileId")
    REFERENCES "File" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "FileInformation"
    ADD CONSTRAINT "FileInformationPublicInfo" FOREIGN KEY ("PublicInfoId")
    REFERENCES "PublicInfo" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "FileInformation"
    ADD CONSTRAINT "FileInformationTest" FOREIGN KEY ("TestId")
    REFERENCES "Test" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "FileInformation"
    ADD CONSTRAINT "FileInformationQuestion" FOREIGN KEY ("QuestionId")
    REFERENCES "Question" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "PublicInfo"
    ADD CONSTRAINT "PublicInfoClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "TestResults"
    ADD CONSTRAINT "TestResultsCourse" FOREIGN KEY ("CourseId")
    REFERENCES "Course" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "TestResults"
    ADD CONSTRAINT "TestResultsTest" FOREIGN KEY ("TestId")
    REFERENCES "Test" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "TestResults"
    ADD CONSTRAINT "TestResultsClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "ActiveTest"
    ADD CONSTRAINT "ActiveTestTest" FOREIGN KEY ("TestId")
    REFERENCES "Test" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "ActiveTest"
    ADD CONSTRAINT "ActiveTestClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "ActiveTest"
    ADD CONSTRAINT "ActiveTestTestType" FOREIGN KEY ("TestTypeId")
    REFERENCES "TestType" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "Question"
    ADD CONSTRAINT "QuestionTest" FOREIGN KEY ("TestId")
    REFERENCES "Test" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "Question"
    ADD CONSTRAINT "QuestionQuestionType" FOREIGN KEY ("QuestionTypeId")
    REFERENCES "QuestionType" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "QuestionResult"
    ADD CONSTRAINT "QuestionResultQuestion" FOREIGN KEY ("QuestionId")
    REFERENCES "Question" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

ALTER TABLE IF EXISTS "QuestionResult"
    ADD CONSTRAINT "QuestionResultClient" FOREIGN KEY ("ClientId")
    REFERENCES "Client" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;


ALTER TABLE IF EXISTS "QuestionType"
    ADD CONSTRAINT "QuestionTypeInputType" FOREIGN KEY ("InputTypeId")
    REFERENCES "InputType" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION;

--/////////////////// Данные

INSERT INTO "Department" ("Title") VALUES ('АиВТ');
INSERT INTO "Department" ("Title") VALUES ('Дизайн');
INSERT INTO "Department" ("Title") VALUES ('Автосервис');

INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('Учителя', 'Учитель', 1);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('269', '269', 1);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('362', '362', 1);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('165', '165', 1);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('Учителя', 'Учитель', 2);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('135', '135', 2);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('325', '325', 2);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('262', '262', 2);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('469', '469', 2);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('Учителя', 'Учитель', 3);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('269', '269', 3);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('161', '161', 3);
INSERT INTO "Group" ("Title", "TitleSingular", "DepartmentId") VALUES ('541', '541', 3);

INSERT INTO "Client" (
    "FirstName",
    "LastName",
    "Patronymic",
    "Password",
    "Email",
    "Telephone",
    "EmailNotifications",
    "ManageCourses",
    "ManageUsers",
    "AploadFiles",
    "ViewYourResults",
    "ViewOtherResults",
    "CreatorId",
    "DepartmentId",
    "GroupId"
) VALUES ('Сергей', 'Молочков', 'Алексеевич', 'asdghsjdjsaqrg345grwfweg235f2', 'molotchkow@mail.ru', 
          '89091784716', false, true, true, true, true, true, NULL, 1, 1);
INSERT INTO "Client" (
    "FirstName",
    "LastName",
    "Patronymic",
    "Password",
    "Email",
    "Telephone",
    "EmailNotifications",
    "ManageCourses",
    "ManageUsers",
    "AploadFiles",
    "ViewYourResults",
    "ViewOtherResults",
    "CreatorId",
    "DepartmentId",
    "GroupId"
) VALUES ('Илья', 'Грязнов', 'Владиславович', '34h5j435kj3jh53hj45jhv346', 'graznow@mail.ru', 
          '89095863712', false, true, true, true, true, true, NULL, 1, 1);
INSERT INTO "Client" (
    "FirstName",
    "LastName",
    "Patronymic",
    "Password",
    "Email",
    "Telephone",
    "EmailNotifications",
    "ManageCourses",
    "ManageUsers",
    "AploadFiles",
    "ViewYourResults",
    "ViewOtherResults",
    "CreatorId",
    "DepartmentId",
    "GroupId"
) VALUES ('Кирилл', 'Полетаев', 'Сергеевич', '3wjb32kb5b2k52bk5k2k3kj', 'poletaew@mail.ru', 
          '89096946146', true, true, true, true, true, true, NULL, 1, 1);

INSERT INTO "Client" (
    "FirstName",
    "LastName",
    "Patronymic",
    "Password",
    "Email",
    "Telephone",
    "EmailNotifications",
    "ManageCourses",
    "ManageUsers",
    "AploadFiles",
    "ViewYourResults",
    "ViewOtherResults",
    "CreatorId",
    "DepartmentId",
    "GroupId"
) VALUES ('Кирилл', 'Зуев', 'Сергеевич', 'dfgdfgdfg324wgw2', 'zuew@mail.ru', 
          '89098496141', false, false, false, false, false, false, NULL, 1, 2);

INSERT INTO "Client" (
    "FirstName",
    "LastName",
    "Patronymic",
    "Password",
    "Email",
    "Telephone",
    "EmailNotifications",
    "ManageCourses",
    "ManageUsers",
    "AploadFiles",
    "ViewYourResults",
    "ViewOtherResults",
    "CreatorId",
    "DepartmentId",
    "GroupId"
) VALUES ('Максим', 'Волков', 'Васильевич', 'j34hb534h5j34b63hjv6j', 'volkov@mail.ru', 
          '8909683735', false, false, false, false, false, false, NULL, 1, 2);