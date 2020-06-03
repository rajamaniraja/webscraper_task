# webscraper_task
![GitHub Logo](/arch.png)



# Command to Run the Application
   docker-compose build
   docker-compose up


# API for scrap will be http://localhost:8001/api/v1/scrap 

# Sample Payload 
{
    "url":"https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/"
 }


# Response Sample 
{
    "name": "PlayStation 4 Pro 1TB Console",
    "description": "Edition:Pro 1TB       PS4 Pro 4K TV GAMING & MORE The most advanced PlayStation system ever. PS4 Pro is designed to take your favorite PS4 games and add to them with more power for graphics, performance, or features for your 4K HDR TV, or 1080p HD TV. Ready to level up?  4K TV Gaming – PS4 Pro outputs gameplay to your 4K TV. Many games, like Call of Duty: WWII, Gran Turismo Sport, and more, are optimized to look stunningly sharp and detailed when played on a 4K TV with PS4 Pro. More HD Power – Turn on Boost Mode to give PS4 games access to the increased power of PS4 Pro. For HD TV Enhanced games, players can benefit from increased image clarity, faster frame rates, or more. HDR Technology – With an HDR TV, compatible PS4 games display an unbelievably vibrant and lifelike range of colors. 4K Entertainment – Stream 4K videos, movies, and shows to your PS4 Pro.*  GREATNESS AWAITS *4K Entertainment requires access to a 4K compatible content streaming service, a robust internet connection, and a compatible 4K display. Enhanced for PS4 Pro Many of the biggest and best PS4 games get an additional boost from PS4 Pro enhancements that fine tune the game’s performance. From the stunning Manhattan skyline of Marvel’s Spider Man and the towering Norse mountains of God of War, to the vast plains of Red Dead Redemption 2 and the intense battlegrounds of Call of Duty: Black Ops 4, you’ll feel the power of your games unleashed wherever you see the Enhanced for PS4 Pro badge**.",
    "reviews": 4940,
    "price": 377,
    "image_url": "{\"https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_.jpg\":[436,500],\"https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_SX466_.jpg\":[406,466],\"https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_SX425_.jpg\":[371,425],\"https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_SX450_.jpg\":[392,450],\"https://images-na.ssl-images-amazon.com/images/I/41GGPRqTZtL._AC_SX355_.jpg\":[310,355]}"
}


