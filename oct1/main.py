from dataclasses import dataclass
import sys
import random
# from typing import Literal

SUITS = ["Hearts", "Diamonds", "Spades", "Clubs"]
FACES = ["Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"]

@dataclass
class Card:
   suit: str
   face: str

@dataclass
class Straight:
   hand: list[Card]

def gen_straight(combs: list[Straight], n: int):
    faces_ex = FACES[:]
    faces_ex.append("Ace")
    i = 0
    while i + n <= 14:
      seq = faces_ex[i:i+n]
      # s = Straight(seq)

      to_r = seq[:]
      to_r.reverse()
      # r = Straight(to_r)
      # print(s)
      # print(r)
      
      combs.append(seq)
      combs.append(to_r)
      print(combs)
      i += 1
      
      

def fill_deck(deck: list[Card]):
   for s in SUITS:
      for r in FACES:
         c = Card(s, r)
         deck.append(c)

def form_hand(n: int, deck: list[Card], comb : list[Straight]) -> bool:
  hand = deck[:n]
  faces = []
  for j in range(len(hand)):
     print(f"{j}: {hand[j].face} of {hand[j].suit}")
     faces.append(hand[j].face)
  
  for i in range(n): deck.pop(0) # modifies array object in place
  #  print(hand)

  # check if straight
  if faces in comb:
    suit = hand[0].suit

    # check if flush
    for s in hand:
       if s.suit != suit:
          return False
    return True
  else:
     return False

if __name__ == '__main__':
  args = sys.argv
  N = int(args[1])
  deck = []
  straight_comb = []

  gen_straight(straight_comb, N)
  fill_deck(deck)
  random.shuffle(deck)

  # print(len(deck), N)
  # print(deck[:N])

  straight_flush = False
  c = 1
  d = 1
  while not straight_flush:
    print(f"Attempt {c}, deck {d}")
    if len(deck) < N:
      fill_deck(deck)
      random.shuffle(deck)
      d += 1
    straight_flush = form_hand(N, deck, straight_comb)
    c += 1
  # print(len(deck))


        

