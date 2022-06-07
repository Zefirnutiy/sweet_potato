CREATE TABLE "Student" (
                  "id" serial NOT NULL,
                  "first_name" varchar(40) NOT NULL,
                  "sure_name" varchar(40) NOT NULL,
                  "email" VARCHAR(120) NOT NULL UNIQUE,
                  "password" VARCHAR(255) NOT NULL,
                  "major" VARCHAR(255) NOT NULL,
                  "course" INTEGER NOT NULL,
                  "group" INTEGER NOT NULL,
                  CONSTRAINT "Student_pk" PRIMARY KEY ("id")
);




CREATE TABLE "Question" (
                  "id" serial NOT NULL,
                  "test_id" INTEGER NOT NULL,
                  "text" VARCHAR(255) NOT NULL,
                  "Answer" INTEGER NOT NULL,
                  "files" INTEGER NOT NULL,
                  "scores" INTEGER NOT NULL,
                  CONSTRAINT "Question_pk" PRIMARY KEY ("id")
);


CREATE TABLE "Question_result" (
                  "id" serial NOT NULL,
                  "student_id" INTEGER NOT NULL,
                  "question_id" INTEGER NOT NULL,
                  "accuracy" BOOLEAN NOT NULL, 
                  CONSTRAINT "Question_result_pk" PRIMARY KEY ("id") 
);


CREATE TABLE "Answer_option" (
                  "id" serial NOT NULL,
                  "question_id" INTEGER NOT NULL,
                  "value" VARCHAR(250) NOT NULL,
                  "true_answer" BOOLEAN NOT NULL,
                  CONSTRAINT "Answer_option_pk" PRIMARY KEY("id")
);


CREATE TABLE "Question_file" (
                  "id" serial NOT NULL,
                  "question_id" INTEGER NOT NULL,
                  "file" VARCHAR(250) NOT NULL,
                  CONSTRAINT "question_file_pk" PRIMARY KEY("id")
);


CREATE TABLE "Answer" (
                  "id" serial NOT NULL,
                  "question_id" integer NOT NULL,
                  "text" VARCHAR(255) NOT NULL,
                  CONSTRAINT "Answer_pk" PRIMARY KEY ("id")
);


CREATE TABLE "Test" (
                  "id" serial NOT NULL,
                  "professor_id" INTEGER NOT NULL,
                  "title" VARCHAR(255) NOT NULL,
                  "start" timestamp NOT NULL,
                  "end" timestamp NOT NULL,
                  "questions_id" INTEGER NOT NULL,
                  "attempt" INTEGER NOT NULL, 
                  "time" TIME NOT NULL,
                  CONSTRAINT "Test_pk" PRIMARY KEY ("id")
);


CREATE TABLE "Result" (
                  "id" serial NOT NULL,
                  "student_id" INTEGER NOT NULL,
                  "test_id" INTEGER NOT NULL,
                  "attempt" INTEGER NOT NULL,
                  "time_spent" TIME NOT NULL,
                  "score_gained" INTEGER NOT NULL,  
                  CONSTRAINT "Result_pk" PRIMARY KEY ("id")
);


CREATE TABLE "Grade" (
                  "id" serial NOT NULL,
                  "test_id" integer NOT NULL,
                  "student_id" integer NOT NULL,
                  CONSTRAINT "Grade_pk" PRIMARY KEY ("id")
);




CREATE TABLE "Additional_materials" (
                  "id" serial NOT NULL,
                  "professor_id" integer NOT NULL,
                  "file" varchar(200) NOT NULL,
                  CONSTRAINT "Additional_materials_pk" PRIMARY KEY ("id")
);




CREATE TABLE "Professor" (
                        "id" serial NOT NULL,
                        "first_name" varchar(40) NOT NULL,
                        "sure_name" varchar(40) NOT NULL,
                        "patronymic" varchar(40) NOT NULL,
                        "email" VARCHAR(120) NOT NULL UNIQUE,
                        "password" VARCHAR(255) NOT NULL,
                        CONSTRAINT "Professor_pk" PRIMARY KEY ("id")
);



ALTER TABLE "Answer" ADD CONSTRAINT "Answer_pk0" FOREIGN KEY ("question_id") REFERENCES "Question"("id");


ALTER TABLE "Test" ADD CONSTRAINT "Test_fk0" FOREIGN KEY ("questions_id") REFERENCES "Question"("id");
ALTER TABLE "Test" ADD CONSTRAINT "Test_fk1" FOREIGN KEY ("professor_id") REFERENCES "Professor"("id");

ALTER TABLE "Question_result" ADD CONSTRAINT "Question_result_fk0" FOREIGN KEY("student_id") REFERENCES "Student"("id");
ALTER TABLE "Question_result" ADD CONSTRAINT "Question_result_fk1" FOREIGN KEY("question_id") REFERENCES "Professor"("id");


ALTER TABLE "Result" ADD CONSTRAINT "Result_fk0" FOREIGN KEY("student_id") REFERENCES "Student"("id");
ALTER TABLE "Result" ADD CONSTRAINT "Result_fk1" FOREIGN KEY("test_id") REFERENCES "Test"("id");


ALTER TABLE "Grade" ADD CONSTRAINT "Grade_fk0" FOREIGN KEY ("test_id") REFERENCES "Test"("id");
ALTER TABLE "Grade" ADD CONSTRAINT "Grade_fk1" FOREIGN KEY ("student_id") REFERENCES "Student"("id");

ALTER TABLE "Additional_materials" ADD CONSTRAINT "Additional_materials_fk0" FOREIGN KEY ("professor_id") REFERENCES "Professor"("id");


ALTER TABLE "Answer_option" ADD CONSTRAINT "Answer_option_fk0" FOREIGN KEY ("question_id") REFERENCES "Question"("id");


ALTER TABLE "question_file" ADD CONSTRAINT "question_file_fk0" FOREIGN KEY ("question_id") REFERENCES "Question"("id");