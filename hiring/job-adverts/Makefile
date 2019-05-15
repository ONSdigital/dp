all: hiring-process job-adverts

hiring-process: backend-engineer-hiring-process frontend-engineer-hiring-process site-reliability-engineer-hiring-process platform-engineer-hiring-process full-stack-engineer-hiring-process

backend-engineer-hiring-process:
	relaxed hiring-process.pug output/backend-engineer-hiring-process.pdf --build-once --locals '{"title":"Backend Engineer", "role": "backend-engineer", "category": "engineering"}'

frontend-engineer-hiring-process:
	relaxed hiring-process.pug output/frontend-engineer-hiring-process.pdf --build-once --locals '{"title":"Frontend Engineer", "role": "frontend-engineer", "category": "engineering"}'

site-reliability-engineer-hiring-process:
	relaxed hiring-process.pug output/site-reliability-engineer-hiring-process.pdf --build-once --locals '{"title":"Site Reliability Engineer", "role": "site-reliability-engineer", "category": "engineering"}'

platform-engineer-hiring-process:
	relaxed hiring-process.pug output/platform-engineer-hiring-process.pdf --build-once --locals '{"title":"Platform Engineer", "role": "platform-engineer", "category": "engineering"}'

full-stack-engineer-hiring-process:
	relaxed hiring-process.pug output/full-stack-engineer-hiring-process.pdf --build-once --locals '{"title":"Full Stack Engineer", "role": "full-stack-engineer", "category": "engineering"}'

job-adverts: backend-engineer-job-advert frontend-engineer-job-advert site-reliability-engineer-job-advert platform-engineer-job-advert full-stack-engineer-job-advert technical-lead-job-advert engineering-manager-job-advert

backend-engineer-job-advert:
	relaxed job-advert.pug output/backend-engineer-job-advert.pdf --build-once --locals '{"title":"Backend Engineer", "role": "backend-engineer", "category": "engineering"}'

frontend-engineer-job-advert:
	relaxed job-advert.pug output/frontend-engineer-job-advert.pdf --build-once --locals '{"title":"Frontend Engineer", "role": "frontend-engineer", "category": "engineering"}'

site-reliability-engineer-job-advert:
	relaxed job-advert.pug output/site-reliability-engineer-job-advert.pdf --build-once --locals '{"title":"Site Reliability Engineer", "role": "site-reliability-engineer", "category": "engineering"}'

platform-engineer-job-advert:
	relaxed job-advert.pug output/platform-engineer-job-advert.pdf --build-once --locals '{"title":"Platform Engineer", "role": "platform-engineer", "category": "engineering"}'

full-stack-engineer-job-advert:
	relaxed job-advert.pug output/full-stack-engineer-job-advert.pdf --build-once --locals '{"title":"Full Stack Engineer", "role": "full-stack-engineer", "category": "engineering"}'

technical-lead-job-advert:
	relaxed job-advert.pug output/technical-lead-job-advert.pdf --build-once --locals '{"title":"Technical Lead", "role": "technical-lead", "category": "leadership"}'

engineering-manager-job-advert:
	relaxed job-advert.pug output/engineering-manager-job-advert.pdf --build-once --locals '{"title":"Engineering Manager", "role": "engineering-manager", "category": "leadership"}'

additional-information:
	relaxed additional-information.pug output/additional-information-original-benefits.pdf --build-once --locals '{"version":"original", "benefits": true}'
	relaxed additional-information.pug output/additional-information-simplified-benefits.pdf --build-once --locals '{"version":"simplified", "benefits": true}'
	relaxed additional-information.pug output/additional-information-original.pdf --build-once --locals '{"version":"original", "benefits": false}'
	relaxed additional-information.pug output/additional-information-simplified.pdf --build-once --locals '{"version":"simplified", "benefits": false}'

.PHONY: all hiring-process backend-engineer-hiring-process full-stack-engineer-hiring-process job-adverts backend-engineer-job-advert frontend-engineer-job-advert site-reliability-engineer-job-advert full-stack-engineer-job-advert technical-lead-job-advert engineering-manager-job-advert
