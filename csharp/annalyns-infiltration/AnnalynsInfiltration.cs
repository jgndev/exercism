using System;

static class QuestLogic
{
    // CanFastAttack can be executed only when the knight is sleeping.
    public static bool CanFastAttack(bool knightIsAwake)
    {
        return !knightIsAwake;
    }

    // CanSpy can be executed if at least one of the characters is awake.
    public static bool CanSpy(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake)
    {
        return knightIsAwake || prisonerIsAwake || archerIsAwake;
    }

    // CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
    public static bool CanSignalPrisoner(bool archerIsAwake, bool prisonerIsAwake)
    {
        return prisonerIsAwake && !archerIsAwake;
    }

    // CanFreePrisoner can be executed if the prisoner is awake and the other two characters are asleep.
    // If Annalyn's pet dog is with her and the archer is sleeping, also returns true.
    public static bool CanFreePrisoner(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake, bool petDogIsPresent)
    {
        return (petDogIsPresent && !archerIsAwake) || prisonerIsAwake && !archerIsAwake && !knightIsAwake;
    }
}
