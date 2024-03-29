class AnnalynsInfiltration {

    // CanFastAttack can be executed only when the knight is sleeping.
    public static boolean canFastAttack(boolean knightIsAwake) {
        return !knightIsAwake;
    }

    // CanSpy can be executed if at least one of the characters is awake.
    public static boolean canSpy(boolean knightIsAwake, boolean archerIsAwake, boolean prisonerIsAwake) {
        return knightIsAwake || archerIsAwake || prisonerIsAwake;
    }

    // CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
    public static boolean canSignalPrisoner(boolean archerIsAwake, boolean prisonerIsAwake) {
        return prisonerIsAwake && !archerIsAwake;
    }

    // CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
    // or if Annalyn's pet dog is with her and the archer is sleeping.
    public static boolean canFreePrisoner(boolean knightIsAwake, boolean archerIsAwake, boolean prisonerIsAwake, boolean petDogIsPresent) {
        return (petDogIsPresent && !archerIsAwake) || (prisonerIsAwake && !knightIsAwake && !archerIsAwake);
    }
}
