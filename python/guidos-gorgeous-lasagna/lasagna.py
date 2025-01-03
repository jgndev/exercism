"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""

EXPECTED_BAKE_TIME = 40
MINUTES_PER_LAYER = 2


def bake_time_remaining(elapsed_bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - elapsed_bake_time


def preparation_time_in_minutes(layers):
    """Calculate preparation time based on lasagna layers.

    :param layers: int - the number of layers for the lasagna.
    :return: int - minutes required for the number of layers.
    """
    return MINUTES_PER_LAYER * layers


def elapsed_time_in_minutes(layers, cook_time):
    """Elapsed cooking time in minutes

    :param layers: int - the number of layers for the lasagna.
    :param cook_time: int - the actual cook_time.
    :return:
    """
    return preparation_time_in_minutes(layers) + cook_time
