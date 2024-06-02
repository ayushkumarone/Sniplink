# Sniplink

This Project is made by Ayush Kumar.

● Shoten Your URLs
  ○ /shorten    -   For shortening your URLs you just need to pass this with the required Shorten ID and Your URL.
                    Users can shorten 5 URLs without registering into our platform and can shorten upto 50 after logging in our platform.

● Access your shortened URLs
  ○ /link/id    -   For getting redirected to your URLs you just need to hit this endpoint where id is the short ID.

● Register
  ○ /register   -   For creating your account you need to hit this endpoint which will create your account.
                    (P.S. - You access all the member exclusive features you need to be a member.)

● Login
  ○ /login      -   For accessing your account you need to hit this endpoint which will generate your API key which would be used for all member exclusive function.
                    (P.S. - This is a member only function so you need to register first.)

● Get analytics about your URLs
  ○ /analytics  -   For getting analytics about your URLs you just need to pass this with your api_key generated.
                    (P.S. - This is a member only function so you need to login before creating your URL to have access to this.)
                    Users can get information about the number of hits, When was URL last hit to monitor their traffic.

● Delete your URL
  ○ /delete/id  -   For deleting your URL you just need to pass this with your api_key generated.
                    (P.S. - This is a member only function so you need to login before creating your URL to have access to this.)
                    
● The links and API keys expire after 24 hours.

