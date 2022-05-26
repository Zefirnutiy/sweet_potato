CREATE TABLE "Student" (
                           "id" serial NOT NULL,
                           "first_name" varchar(40) NOT NULL,
                           "last_name" varchar(40) NOT NULL,
                           "email" VARCHAR(120) NOT NULL UNIQUE,
                           "password" VARCHAR(255) NOT NULL,
                           "n_group" varchar(5) NOT NULL,
                           CONSTRAINT "Student_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "Question" (
                            "id" serial NOT NULL,
                            "title" VARCHAR(255) NOT NULL,
                            "answer" VARCHAR(255) NOT NULL,
                            CONSTRAINT "Question_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );

CREATE TABLE "Answers" (
                          "id" serial NOT NULL,
                          "question_id" integer NOT NULL,
                          "text" VARCHAR(255) NOT NULL,
                          CONSTRAINT "Answers_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );

CREATE TABLE "Test" (
                        "id" serial NOT NULL,
                        "title" VARCHAR(255) NOT NULL,
                        "questions" integer NOT NULL,
                        "date_start" timestamp NOT NULL,
                        "date_end" timestamp NOT NULL,
                        CONSTRAINT "Test_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "Grade" (
                         "id" serial NOT NULL,
                         "test" integer NOT NULL,
                         "student" integer NOT NULL,
                         CONSTRAINT "Grade_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "Additional_materials" (
                                        "id" serial NOT NULL,
                                        "file" varchar(200) NOT NULL,
                                        "test" integer NOT NULL,
                                        CONSTRAINT "Additional_materials_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "Professor" (
                             "id" serial NOT NULL,
                             "first_name" varchar(40) NOT NULL,
                             "last_name" varchar(40) NOT NULL,
                             "patronymic" varchar(40) NOT NULL,
                             "email" VARCHAR(120) NOT NULL UNIQUE,
                             "password" VARCHAR(255) NOT NULL,
                             "tests" integer NOT NULL,
                             CONSTRAINT "Professor_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );





ALTER TABLE "Answers" ADD CONSTRAINT "Answers_pk0" FOREIGN KEY ("question_id") REFERENCES "Question"("id");

ALTER TABLE "Test" ADD CONSTRAINT "Test_fk0" FOREIGN KEY ("questions") REFERENCES "Question"("id");


ALTER TABLE "Grade" ADD CONSTRAINT "Grade_fk0" FOREIGN KEY ("test") REFERENCES "Test"("id");
ALTER TABLE "Grade" ADD CONSTRAINT "Grade_fk1" FOREIGN KEY ("student") REFERENCES "Student"("id");

ALTER TABLE "Additional_materials" ADD CONSTRAINT "Additional_materials_fk0" FOREIGN KEY ("test") REFERENCES "Test"("id");

ALTER TABLE "Professor" ADD CONSTRAINT "Professor_fk0" FOREIGN KEY ("tests") REFERENCES "Test"("id");