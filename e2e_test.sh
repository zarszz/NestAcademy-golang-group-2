psql postgresql://postgres:password@localhost:5432/nestacademy_golang -c "DROP SCHEMA public CASCADE" &&
psql postgresql://postgres:password@localhost:5432/nestacademy_golang -c "CREATE SCHEMA public" &&

(go run ./main.go &) &&

sleep 120 &&

psql postgresql://postgres:password@localhost:5432/nestacademy_golang -f ./insert_admin.sql &&


newman run final_project_go.postman_collection.json -e final_project_go.postman_environment.json &&

psql postgresql://postgres:password@localhost:5432/nestacademy_golang -c "DELETE FROM public.users WHERE users.email = 'admin@gmail.com'" &&

fuser -k 4444/tcp