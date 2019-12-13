1. Go to BotFather in Telegram and create the bot
2. Set to your bot three commands:
    
    /info - Show server info
    
    /mute - Disable notification about users connections
    
    /mute - Enable notification about users connections

3. Set you bot token in to .env/BOT_TOKEN
4. If you want configure proxy for telegram connection, set .env/PROXY_*
5. Set in .env/ADMIN_PASSWORD for allow execute rcon command over bot on game server.
   Rcon command will be executed only from direct messages.
6. Set rcon password for game server in .env/RCON_PASSWORD 

For start `./start.sh`

Run in counter strike console command `connect 127.0.0.1:27016`

For stop  `./stop.sh`


 