all: how-we-slack your-first-few-weeks useful-things-to-know

how-we-slack:
	relaxed how-we-slack.pug output/how-we-slack.pdf --build-once --locals '{"team": "development", "role": "frontend-engineer"}'

your-first-few-weeks:
	relaxed your-first-few-weeks.pug output/your-first-few-weeks.pdf --build-once --locals '{"name": "Somebody", "team": "development", "role": "frontend-engineer"}'

useful-things-to-know:
	relaxed useful-things-to-know.pug output/useful-things-to-know.pdf --build-once

.PHONY: all how-we-slack your-first-few-weeks useful-things-to-know
