CREATE TABLE "Admin" (
    "Id" serial  NOT NULL,
    "FirstName" varchar(20)  NOT NULL,
    "LastName" varchar(30)  NOT NULL,
    "Password" varchar(60)  NOT NULL,
    "Email" varchar(256)  NULL,
    CONSTRAINT "AdminPk" PRIMARY KEY ("Id")
);

CREATE TABLE "Level" (
    "Id" smallserial  NOT NULL,
    "Title" varchar(100)  NOT NULL,
    "Price" int  NOT NULL,
    "Paid" boolean  default false,
    "CreateCourse" boolean  default false,
    "TakeCourse" boolean  default false,
    "AploadFile" boolean  default false,
    "ViewYourResult" boolean  default false,
    "ViewOtherResults" boolean  default false,
    CONSTRAINT "LevelPk" PRIMARY KEY ("Id")
);

CREATE TABLE "Organization" (
    "Id" serial  NOT NULL,
    "Title" varchar(50)  NOT NULL,
    "Password" varchar(60)  NOT NULL,
    "Email" varchar(256)  NULL,
    "EmailNotifications" boolean  default false,
    "LevelId" smallint  NOT null default 0,
    CONSTRAINT "OrganizationPk" PRIMARY KEY ("Id")
);

CREATE TABLE "DeadLine" (
    "Id" serial  NOT NULL,
    "Date" timestamp  NOT NULL,
    "LevelId" smallint  NOT NULL,
    "OrganizationId" int8  NOT NULL,
    CONSTRAINT "DeadLinePk" PRIMARY KEY ("Id")
);

CREATE TABLE "Payment" (
    "Number" int  NOT NULL,
    "Name" varchar(100)  NOT NULL,
    "Money" int  NOT NULL,
    "Date" timestamp  NOT NULL,
    "LevelId" smallint  NOT NULL,
    "OrganizationId" int8  NOT NULL,
    CONSTRAINT "PaymentPk" PRIMARY KEY ("Number")
);

ALTER TABLE "DeadLine" ADD CONSTRAINT "OrganisationOrganisationLevel"
    FOREIGN KEY ("LevelId")
    REFERENCES "Level" ("Id")  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Payment" ADD CONSTRAINT "PaymentOrganization"
    FOREIGN KEY ("OrganizationId")
    REFERENCES "Organization" ("Id")  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "Payment" ADD CONSTRAINT "PaymentLevel"
    FOREIGN KEY ("LevelId")
    REFERENCES "Level" ("Id")  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

ALTER TABLE "DeadLine" ADD CONSTRAINT "DeadLineOrganization"
    FOREIGN KEY ("OrganizationId")
    REFERENCES "Organization" ("Id")  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;
