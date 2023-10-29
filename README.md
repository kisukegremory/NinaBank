# NinaBank

This projects seeks to build a fake bank webapp, where the user after its authentication will be able to:
- bank balance
- bank deposit (fake)
- cash withdrawal
- cash transfer between ninaccounts

The interface will be implemented using htmx, we will use golang with gin and sqlc/sqlx to be our server side render and backend, kafka will be used for message queues, mysql will be used as our database, those choices were made to learn some of these tecnologies that i never been able to use in my last projects, in specific htmx, kafka and sqlc


These project has the purpose of being an opportunity to design a scalable system that could upscaled horinzontally, the only point that is not been considered here is the database reliability, if we were going to a production enviroment i would pick the AWS RDS to manage the database and apply the backups and replicas within different regions to be sure that will be persisted, but for our study purposes we are going to rely on docker images!

