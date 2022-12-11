from django.contrib import admin

from app.models import Action, Match, Player, Team, MyTeam

# Register your models here.
admin.site.register(Player)
admin.site.register(Team)
admin.site.register(MyTeam)
admin.site.register(Match)
admin.site.register(Action)