# Short Link Kata
A server that acts as a bit.ly-like URL shortener. The primary interface should is a JSON API that allows the following:
* Create a random short link for arbitrary URLs, e.g., bit.ly/2FhfhXh
* The same URL should always generate the same random shortlink
* Allow creating custom short links to arbitrary URLs, e.g., bit.ly/my-custom-link
* Provide a route for returning stats in a given short link, including:
    - When the short link was created
    - How many times the short link has been visited total
    - A histogram of number of visits to the short link per day
* The server itself handles redirecting short links to the URLs it creates

# Configuration

I used PostgreSQL version 14.1 as my database.  For simplicity, I did not configure a password (though I would in a production environment).

To setup the database, please execute the commands (in the specified order) in the db/schema.sql file.


# Next steps

I'd like to:
* backfill the tests that I ended up testing manually due to deadline constraints
* make the server robust against bad input (validating URLs)
* use the database data to create a histogram image
* enable access to histogram images by API
* extract the code in db.SlugVisited function so that each operation is its own function
* implement database connection pooling to increase the speed of the redirect API
