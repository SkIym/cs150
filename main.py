from dataclasses import dataclass, field
import random
from enum import Enum

class Face(Enum):
    Ace= 1
    Two = 2
    Three = 3
    Four = 4
    Five = 5
    Six = 6
    Seven = 7
    Eight = 8
    Nine = 9
    Ten = 10
    Jack = 11
    Queen = 12
    King = 13

suits = ["Hearts", "Diamonds", "Spades", "Clubs"]

@dataclass
class Card:
    suit: str
    face: Face

# @dataclass
# class PlayerDeck:
#   cards: list[Card] = field(default_factory=list)

@dataclass
class Deck:
    cards: list[Card] = field(default_factory=list)

    def fill(self):
        for s in suits:
            for f in Face:
                c = Card(s, f)
                self.cards.append(c)
  
    def shuffle(self):
        random.shuffle(self.cards)
  
    def split(self) -> tuple[list[Card], list[Card]]:
        first = self.cards[:26]
        second = self.cards[26:]
        return first, second

@dataclass
class Player:
    name: str = field(default="Player")
    cards: list[Card] = field(default_factory=list)
    hand: list[Card] = field(default_factory=list)

    def draw(self) -> Card:
        card = self.cards.pop()
        self.hand.append(card)
        return card
  
    def draw_two(self) -> tuple[Card, Card]:
        face_up = self.cards.pop()
        face_down = self.cards.pop()
        self.hand.append(face_down)
        self.hand.append(face_up)
        return face_up, face_down

    def place(self, drawn: list[Card]):
        random.shuffle(drawn)
        for card in drawn:
            self.cards.insert(0, card)
  
@dataclass
class WarView:
    player1_name: str = field(default="Player")
    player2_name: str = field(default="Player")

    def print_draw_result(self, player1_card: Card, player2_card: Card):
        print(f"{self.player1_name}:")
        print(f"- {player1_card.face.name} of {player1_card.suit}")
        print('')
        print(f"{self.player2_name}:")
        print(f"- {player2_card.face.name} of {player2_card.suit}")
        print('')

    def print_war_draw_result(self, player1_fup: Card, player1_fdown: Card, player2_fup: Card, player2_fdown: Card):
        print(f"{self.player1_name}:")
        print(f"- {player1_fup.face.name} of {player1_fup.suit}")
        print(f"- {player1_fdown.face.name} of {player1_fdown.suit} (face down)")
        print('')
        print(f"{self.player2_name}:")
        print(f"- {player2_fup.face.name} of {player2_fup.suit}")
        print(f"- {player2_fdown.face.name} of {player2_fdown.suit} (face down)")
        print('')

    def print_round_result(self, player_name: str, winning_card: Card):
        print(f"Round winner: {player_name} ({winning_card.face.name} of {winning_card.suit})")
        print('')

    def print_war_notif(self, new: bool):
        if (new):
            print('Commencing war...')
        else:
            print("Continuing war...")
        print('')

    def print_game_result(self, player_name: str, num_cards: int):
        print(f'{player_name} wins with {num_cards} cards in their deck')

    def print_round_number(self, round):
        print(f"Round {round}")
        print('')
    
    def print_war_number(self, round, war):
        print(f"Round {round}, War {war}")
        print('')

@dataclass
class War:
    view: WarView
    player1: Player
    player2: Player
    deck: Deck = Deck()
    round: int = 1
    war: int = 1
    is_game_over: bool = False
    winning_player: Player = Player()
    winning_player_cards_left: int = 0

    def have_cards_left(self) -> bool:
        if len(self.player1.cards) < 2 or len(self.player2.cards) < 2:
            return False
        return True
  
    def clear_hands(self):
        self.player1.hand.clear()
        self.player2.hand.clear()

        # resets war
        self.war = 1
  
    def get_temp_winner(self) -> tuple[Player, Card]:
        if len(self.player2.cards) > len(self.player2.cards):
            return self.player2, self.player2.hand[-1]
        else:
            return self.player1, self.player1.hand[-1]
  
    def get_winner(self, player1_card: Card, player2_card: Card) -> tuple[Player, Card]:
        if player1_card.face.value > player2_card.face.value:
            return self.player1, player1_card
        elif player1_card.face.value < player2_card.face.value:
            return self.player2, player2_card

    def start(self):
        self.deck.fill()
        self.deck.shuffle()
        self.player1.cards, self.player2.cards = self.deck.split()

        while not self.is_game_over:
            self.play()
            # print(len(self.player1.cards), len(self.player2.cards))

        self.view.print_game_result(self.winning_player.name, self.winning_player_cards_left)

    def commence_war(self) -> tuple[Player, Card]:

        ongoing = True
        while ongoing:

            if not self.have_cards_left():
                return self.get_temp_winner()
            
            self.view.print_war_number(self.round, self.war)

            player1_fup, player1_fdown = self.player1.draw_two()
            player2_fup, player2_fdown = self.player2.draw_two()

            self.view.print_war_draw_result(player1_fup, player1_fdown, player2_fup, player2_fdown)

            if player1_fup.face.value != player2_fup.face.value:
                winner, winning_card = self.get_winner(player1_fup, player2_fup)
                ongoing = False
            elif (player1_fup.face.value == player2_fup.face.value) and self.have_cards_left():
                self.view.print_war_notif(False)
                self.war += 1

        return winner, winning_card
  
    def play(self):

        self.view.print_round_number(self.round)
        player1_fup = self.player1.draw()
        player2_fup = self.player2.draw()

        self.view.print_draw_result(player1_fup, player2_fup)
        winner, winning_card = self.get_temp_winner()

        if player1_fup.face.value != player2_fup.face.value:
            winner, winning_card = self.get_winner(player1_fup, player2_fup)
        else:
            if self.have_cards_left():
                self.view.print_war_notif(True)
                winner, winning_card = self.commence_war()

        self.view.print_round_result(winner.name, winning_card)
        winner.place(self.player1.hand + self.player2.hand)
        
        if (len(self.player1.cards) == 0 or len(self.player2.cards) == 0):
            self.is_game_over = True
            self.winning_player = winner
            self.winning_player_cards_left = len(winner.cards) - (len(self.player1.hand) + len(self.player2.hand))
        
        self.clear_hands()
        self.round += 1

if __name__ == '__main__':
    view = WarView("Player 1", "Player 2") # view
    # Deck, Player, and Card classes - model
    game = War(view=view, player1=Player("Player 1"), player2=Player("Player 2")) # controller
    game.start()