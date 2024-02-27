feedback=# SELECT * FROM feedbacks
feedback-# 
feedback-# ;
 id | created_at | updated_at | deleted_at 
----+------------+------------+------------
(0 rows)

feedback=# 

<!-- RESOLVED used auto migrate on Feedback struct but only created columns from gorm.Model super class. Look into this next. -->