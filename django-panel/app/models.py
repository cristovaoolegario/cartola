from django.db import models

# Create your models here.
class Player(models.Model):
  name = models.CharField(max_length=50)
  initial_price = models.FloatField()

  def __str__(self):
    return self.name

class Team(models.Model):
  name = models.CharField(max_length=50)
  
  def __str__(self):
        return self.name

class MyTeam(models.Model):
  players = models.ManyToManyField(Player)

  def __str__(self):
        return [player.name for player in self.players.all()].__str__()

class Match(models.Model):
    match_date = models.DateTimeField()
    team_a = models.ForeignKey(
        Team, on_delete=models.PROTECT, related_name='team_a_matches')
    team_a_goal = models.IntegerField(default=0)
    team_b = models.ForeignKey(
        Team, on_delete=models.PROTECT, related_name='team_b_matches')
    team_b_goal = models.IntegerField(default=0)

    def __str__(self):
        return f"{self.team_a.name} x {self.team_b.name}"