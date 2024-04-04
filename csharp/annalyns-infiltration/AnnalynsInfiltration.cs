using System;

static class QuestLogic
{
    
    // CanFastAttack can be executed only when the knight is sleeping.
    public static bool CanFastAttack(bool knightIsAwake) 
        => !knightIsAwake;

    // CanSpy can be executed if at least one of the characters is awake.
    public static bool CanSpy(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake) 
        => knightIsAwake || archerIsAwake || prisonerIsAwake;

    // CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
    public static bool CanSignalPrisoner(bool archerIsAwake, bool prisonerIsAwake)
        => prisonerIsAwake && !archerIsAwake;

    // CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep,
    // or if the player's pet dog is present and the archer is sleeping.
    public static bool CanFreePrisoner(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake, bool petDogIsPresent)
        => petDogIsPresent && !archerIsAwake || prisonerIsAwake && !knightIsAwake && !archerIsAwake;
}
