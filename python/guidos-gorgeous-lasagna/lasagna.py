"""Functions used in preparing Guido's gorgeous lasagna.

Learn about Guido, the creator of the Python language:
https://en.wikipedia.org/wiki/Guido_van_Rossum

This is a module docstring, used to describe the functionality
of a module and its functions and/or classes.
"""


EXPECTED_BAKE_TIME = 40


def bake_time_remaining(bake_time):
    """Calculate the bake time remaining.

    :param elapsed_bake_time: int - baking time already elapsed.
    :return: int - remaining bake time (in minutes) derived from 'EXPECTED_BAKE_TIME'.

    Function that takes the actual minutes the lasagna has been in the oven as
    an argument and returns how many minutes the lasagna still needs to bake
    based on the `EXPECTED_BAKE_TIME`.
    """

    return EXPECTED_BAKE_TIME - bake_time


def preparation_time_in_minutes(layers):
    """Calculate the number of minutes spent on layers.

    :param layers: int - the number of layers.
    :return: int - the number of layers multiplied by 2, each layer takes 2 minutes.
    """
    return layers * 2



def elapsed_time_in_minutes(layers, cook_time):
    """Calculate the elapsed time spent on the lasagna.

    param: layers: int - the number of layers.
    param: cook_time: int - the number of minutes spent cooking.
    return: the calculated layer preparation time + time spent cooking.
    """
    return preparation_time_in_minutes(layers) + cook_time
