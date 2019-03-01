all: hiring-process job-adverts

hiring-process: backend-engineer-hiring-process frontend-engineer-hiring-process site-reliability-engineer-hiring-process platform-engineer-hiring-process full-stack-engineer-hiring-process

backend-engineer-hiring-process:
	relaxed hiring-process.pug output/backend-engineer-hiring-process.pdf --build-once --locals '{"title":"Backend Engineer", "role": "backend-engineer"}'

frontend-engineer-hiring-process:
	relaxed hiring-process.pug output/frontend-engineer-hiring-process.pdf --build-once --locals '{"title":"Frontend Engineer", "role": "frontend-engineer"}'

site-reliability-engineer-hiring-process:
	relaxed hiring-process.pug output/site-reliability-engineer-hiring-process.pdf --build-once --locals '{"title":"Site Reliability Engineer", "role": "site-reliability-engineer"}'

platform-engineer-hiring-process:
	relaxed hiring-process.pug output/platform-engineer-hiring-process.pdf --build-once --locals '{"title":"Platform Engineer", "role": "platform-engineer"}'

full-stack-engineer-hiring-process:
	relaxed hiring-process.pug output/full-stack-engineer-hiring-process.pdf --build-once --locals '{"title":"Full Stack Engineer", "role": "full-stack-engineer"}'

job-adverts: backend-engineer-job-advert frontend-engineer-job-advert site-reliability-engineer-job-advert platform-engineer-job-advert full-stack-engineer-job-advert

backend-engineer-job-advert:
	relaxed job-advert.pug output/backend-engineer-job-advert.pdf --build-once --locals '{"title":"Backend Engineer", "role": "backend-engineer"}'

frontend-engineer-job-advert:
	relaxed job-advert.pug output/frontend-engineer-job-advert.pdf --build-once --locals '{"title":"Frontend Engineer", "role": "frontend-engineer"}'

site-reliability-engineer-job-advert:
	relaxed job-advert.pug output/site-reliability-engineer-job-advert.pdf --build-once --locals '{"title":"Site Reliability Engineer", "role": "site-reliability-engineer"}'

platform-engineer-job-advert:
	relaxed job-advert.pug output/platform-engineer-job-advert.pdf --build-once --locals '{"title":"Platform Engineer", "role": "platform-engineer"}'

full-stack-engineer-job-advert:
	relaxed job-advert.pug output/full-stack-engineer-job-advert.pdf --build-once --locals '{"title":"Full Stack Engineer", "role": "full-stack-engineer"}'

.PHONY: all hiring-process backend-engineer-hiring-process full-stack-engineer-hiring-process job-adverts backend-engineer-job-advert frontend-engineer-job-advert site-reliability-engineer-job-advert full-stack-engineer-job-advert
