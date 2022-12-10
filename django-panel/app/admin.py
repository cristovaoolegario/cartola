from django.contrib import admin

from app.models import Player, Team, MyTeam

# Register your models here.
admin.site.register(Player)
admin.site.register(Team)
admin.site.register(MyTeam)