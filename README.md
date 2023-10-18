# About project
Simple single page film catalog. It start from 3 pre-added films and provide form to add own films.

What it does:
* Provides form to add film - POST-request sent on submit and handled on backend part (Go)
* Films list can be updated with newly added film - HTMX handles response and updates UI with data from server
* Keeps films "db" in memory, means there is no persistance across sessions
* It uses HTML Templates and HTML Template Fragments for targeted UI updates 

What it does NOT (yet):
* Does not store films in any kind of DB
* Does not validate input - you can submit same form multiple times and same film will be added multiple times; you can even submit empty form.
* Form doesn't clear after adding movie to list

# Purpose
I started learning web development using Go + HTMX and wanted to try to create something that would work. I chose HTMX as the front-end library because of its ease of use. This is a very simple project, the purpose of which is to become more familiar with technology.