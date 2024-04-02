using System;

static class QuestLogic  {
    // CanFastAttack can only be executed when the knight is sleeping.
    public static bool CanFastAttack(bool knightIsAwake) {
        return !knightIsAwake;
    }

    // CanSPy can be executed if at least one of the characters is awake.
    public static bool CanSpy(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake) {
        return knightIsAwake || archerIsAwake || prisonerIsAwake;
    }

    // CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
    public static bool CanSignalPrisoner(bool archerIsAwake, bool prisonerIsAwake) {
        return !archerIsAwake && prisonerIsAwake;
    }

    // CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
    // or if Annalyn's pet dog is with her and the archer is sleeping.
    public static bool CanFreePrisoner(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake, bool petDogIsPresent) {
        if (petDogIsPresent && !archerIsAwake) {
            return true;
        }

        if (prisonerIsAwake && !knightIsAwake && !archerIsAwake) {
            return true;
        }

        return false;
    }
}
