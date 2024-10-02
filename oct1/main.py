from dataclasses import dataclass
import sys
import random
from typing import Literal, TypeVar
from enum import Enum

# i only cloned the repo when i finished
# Start: 2:30 PM
# End: 3:42 PM

SUIT = Literal["Hearts", "Diamonds", "Spades", "Clubs"]
suits = ["Hearts", "Diamonds", "Spades", "Clubs"]
FACE = Literal["Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"]
faces = ["Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"]

@dataclass
class Card:
   suit: SUIT
   face: FACE

@dataclass
class Deck:
  drawpile: list[Card]
  discard: list[Card]
  
  def fill_deck(self):
    for s in suits:
        for r in faces:
          c = Card(s, r)
          self.drawpile.append(c)

  def return_discard(self):
    self.drawpile.extend(self.discard)
    self.discard.clear()
  
  def shuffle(self):
    random.shuffle(drawpile)


def gen_straight(combs: list[list[FACE]], n: int):
    faces_ex = faces[:]
    faces_ex.append("Ace")
    i = 0
    while i + n <= 14:
      seq = faces_ex[i:i+n]
      combs.append(seq)
      i += 1

def form_hand(n: int, deck: list[Card], comb : list[list[FACE]], discard: list[Card]) -> bool:
  hand = deck[:n]
  faces: list[FACE] = []
  for j in range(len(hand)):
     print(f"{j + 1}: {hand[j].face} of {hand[j].suit}")
     faces.append(hand[j].face)
  
  for i in range(n): discard.append(deck.pop(0)) 

  # check if straight
  for c in comb:

     if set(c).difference(set(faces)) == set():
        suit = hand[0].suit
        
        # check if flush
        if len(set([s.suit for s in hand])) == 1:
          return True
  
  return False

if __name__ == '__main__':
  args = sys.argv
  N = int(args[1])
  deck = Deck([], [])
  deck.fill_deck()

  straight_comb = []
  gen_straight(straight_comb, N)

  drawpile = deck.drawpile
  discardpile = deck.discard
  deck.shuffle()

  straight_flush = False
  c = 0
  d = 0

  while not straight_flush:
    c += 1

    if len(drawpile) < N:  
      deck.return_discard()
      deck.shuffle()
      d += 1

    print(f"Attempt {c}, deck {d + 1}")
    straight_flush = form_hand(N, drawpile, straight_comb, discardpile)
    # print(len(drawpile), len(discardpile))
