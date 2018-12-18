all: hiring-process job-adverts

hiring-process:
	relaxed hiring-process.pug output/backend-engineer-hiring-process.pdf --build-once --locals '{"title":"Backend Engineer", "role": "backend-engineer"}'

job-adverts: backend-engineer-job-advert frontend-engineer-job-advert site-reliability-engineer-job-advert full-stack-engineer-job-advert

backend-engineer-job-advert:
	relaxed job-advert.pug output/backend-engineer-job-advert.pdf --build-once --locals '{"title":"Backend Engineer", "role": "backend-engineer"}'

frontend-engineer-job-advert:
	relaxed job-advert.pug output/frontend-engineer-job-advert.pdf --build-once --locals '{"title":"Frontend Engineer", "role": "frontend-engineer"}'

site-reliability-engineer-job-advert:
	relaxed job-advert.pug output/site-reliability-engineer-job-advert.pdf --build-once --locals '{"title":"Site Reliability Engineer", "role": "site-reliability-engineer"}'

full-stack-engineer-job-advert:
	relaxed job-advert.pug output/full-stack-engineer-job-advert.pdf --build-once --locals '{"title":"Full Stack Engineer", "role": "full-stack-engineer"}'

.PHONY: all hiring-process job-adverts backend-engineer-job-advert frontend-engineer-job-advert site-reliability-engineer-job-advert full-stack-engineer-job-advert
