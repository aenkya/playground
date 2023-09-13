## Design a deck of Cards

### Constraints and assumptions
**Is this a generic deck of cards for games like poker and black jack?**

Yes, design a generic deck then extend it to black jack

**Can we assume the deck has 52 cards (2-10, Jack, Queen, King, Ace) and 4 suits?**

Yes

**Can we assume inputs are valid or do we have to validate them?**

Assume they're valid

## Approach

**Terminology**

A standard pack of playing cards consists of 52 cards, divided into four suits: clubs, diamonds, hearts, and spades. Each suit has 13 cards, which range from ace (the highest-ranking card) to 2 (the lowest-ranking card). The face cards in a standard deck are the jack, queen, and king.

The suits are represented by different symbols:

* Clubs: a black club with three parallel lines
* Diamonds: a red diamond with one or more dots
* Hearts: a red heart with one or more dots
* Spades: a black spade with one or more dots

The face cards are depicted as follows:

* Jack: a young man or boy
* Queen: a woman
* King: a man

The ace can be used as either a 1 or an 11, depending on the game being played.

The standard deck of 52 cards is often augmented with two jokers, which are wild cards. Jokers can be used as any card in the deck, and they can also be used to draw additional cards.

The standard deck of 52 cards is used in a wide variety of games, including poker, blackjack, and rummy. It is also a popular choice for card tricks and magic shows.

Here are some of the most popular games that use a standard deck of playing cards:

* Poker: A card game in which players bet on the value of their hands.
* Blackjack: A card game in which players attempt to get as close to 21 as possible without going over.
* Rummy: A card game in which players try to make melds of cards with the same rank or suit.
* Bridge: A card game in which two teams of two players each compete to win tricks.
* Solitaire: A card game in which players try to arrange cards in a specific order.

The standard deck of playing cards is a versatile and popular tool that can be used for a variety of purposes. It is a great way to have fun with friends and family, or to improve your skills at card games.

**Major**

Ugandan cards. The 12 face cards are Jack, Queen and King for every suit(Village). There could be jokers but for the classic we'll leave it out.
The largest card is the Ace. The Ace of the Spades suite is the highest rated ace at 60. The other aces are 15. The rest are numbered from 1 to 10 for each suit. They are worth their number in ratings.


### Example Solution

```python
from abc import ABCMeta, abstractmethod
from enum import Enum
import sys


class Suit(Enum):

    HEART = 0
    DIAMOND = 1
    CLUBS = 2
    SPADE = 3


class Card(metaclass=ABCMeta):

    def __init__(self, value, suit):
        self.value = value
        self.suit = suit
        self.is_available = True

    @property
    @abstractmethod
    def value(self):
        pass

    @value.setter
    @abstractmethod
    def value(self, other):
        pass


class BlackJackCard(Card):

    def __init__(self, value, suit):
        super(BlackJackCard, self).__init__(value, suit)

    def is_ace(self):
        return True if self._value == 1 else False

    def is_face_card(self):
        """Jack = 11, Queen = 12, King = 13"""
        return True if 10 < self._value <= 13 else False

    @property
    def value(self):
        if self.is_ace() == 1:
            return 1
        elif self.is_face_card():
            return 10
        else:
            return self._value

    @value.setter
    def value(self, new_value):
        if 1 <= new_value <= 13:
            self._value = new_value
        else:
            raise ValueError('Invalid card value: {}'.format(new_value))


class Hand(object):

    def __init__(self, cards):
        self.cards = cards

    def add_card(self, card):
        self.cards.append(card)

    def score(self):
        total_value = 0
        for card in card:
            total_value += card.value
        return total_value


class BlackJackHand(Hand):

    BLACKJACK = 21

    def __init__(self, cards):
        super(BlackJackHand, self).__init__(cards)

    def score(self):
        min_over = sys.MAXSIZE
        max_under = -sys.MAXSIZE
        for score in self.possible_scores():
            if self.BLACKJACK < score < min_over:
                min_over = score
            elif max_under < score <= self.BLACKJACK:
                max_under = score
        return max_under if max_under != -sys.MAXSIZE else min_over

    def possible_scores(self):
        """Return a list of possible scores, taking Aces into account."""
        # ...


class Deck(object):

    def __init__(self, cards):
        self.cards = cards
        self.deal_index = 0

    def remaining_cards(self):
        return len(self.cards) - deal_index

    def deal_card():
        try:
            card = self.cards[self.deal_index]
            card.is_available = False
            self.deal_index += 1
        except IndexError:
            return None
        return card

    def shuffle(self):  # ...

```