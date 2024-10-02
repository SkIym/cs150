from dataclasses import dataclass
import sys
import random
from typing import Literal

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
class Hand:
  n: int
  faces: list[FACE]
  suits: list[SUIT]

  def check_straight(self, comb: list[list[FACE]]) -> bool:
    for c in comb:
      if set(c).difference(set(self.faces)) == set():
          return True
      
  def check_flush(self) -> bool:
    if len(set(self.suits)) == 1:
          return True

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
    random.shuffle(self.drawpile)
  
  def draw(self, n: int) -> Hand:
    hand = Hand(n, [],[])
    for i in range(n):
      card = self.drawpile.pop(0)
      hand.faces.append(card.face)
      hand.suits.append(card.suit)
      self.discard.append(card)
    return hand

# generate straight combinations
def gen_straight(combs: list[list[FACE]], n: int):
    faces_ex = faces[:]
    faces_ex.append("Ace")
    i = 0
    while i + n <= 14:
      seq = faces_ex[i:i+n]
      combs.append(seq)
      i += 1

def form_hand(n: int, comb: list[list[FACE]], deck: Deck) -> bool:
  hand = deck.draw(n)

  for j in range(hand.n):
     print(f"{j + 1}: {hand.faces[j]} of {hand.suits[j]}")

  print('')

  # check if straight flush
  if hand.check_flush():
     if hand.check_straight(comb):
        return True
  
  return False

if __name__ == '__main__':
  args = sys.argv
  N = int(args[1])
  deck = Deck([], [])
  deck.fill_deck()

  straight_comb = []
  gen_straight(straight_comb, N)

  deck.shuffle()

  straight_flush = False
  c = 0
  d = 0

  while not straight_flush:
    c += 1

    if len(deck.drawpile) < N:  
      deck.return_discard()
      deck.shuffle()
      d += 1

    print(f"Attempt {c}, deck {d + 1}")
    straight_flush = form_hand(N, straight_comb, deck)
    # print(len(deck.drawpile), len(deck.discard))
