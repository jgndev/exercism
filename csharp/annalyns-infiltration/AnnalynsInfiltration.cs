using System;

static class QuestLogic
{
    // CanFastAttack can be executed if the knight is sleeping.
    public static bool CanFastAttack(bool knightIsAwake) => !knightIsAwake;

    // CanSpy can be executed if at least one of the characters is awake
    public static bool CanSpy(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake)
        => knightIsAwake || archerIsAwake || prisonerIsAwake;

    // CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
    public static bool CanSignalPrisoner(bool archerIsAwake, bool prisonerIsAwake)
        => prisonerIsAwake && !archerIsAwake;

    // CanFreePrisoner can be executed if the prisoner is awake and the other two characters are sleeping
    // or if Annalyn's pet dog is with her and the archer is sleeping.
    public static bool CanFreePrisoner(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake,
        bool petDogIsPresent)
        => (petDogIsPresent && !archerIsAwake) || (prisonerIsAwake && !knightIsAwake && !archerIsAwake);
}
