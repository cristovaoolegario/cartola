from django.test import TestCase
from .models import Match, Team
from datetime import datetime
from django.utils import timezone


# Create your tests here.

class TestMatch(TestCase):

  def test_match_to_string(self):
    team_a = Team.objects.create(name='team_a')
    team_b = Team.objects.create(name='team_b')
    match = Match.objects.create(match_date=datetime.now(tz=timezone.utc),team_a=team_a,team_b=team_b)

    self.assertEqual('team_a x team_b', match.__str__())
    