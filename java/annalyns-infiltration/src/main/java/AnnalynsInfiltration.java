class AnnalynsInfiltration {

    /**
     * Determines if a fast attack can be executed based on the knight's wakefulness.
     *
     * @param knightIsAwake the wakefulness of the knight; {@code true} if the knight is awake, {@code false} otherwise.
     * @return {@code true} if the fast attack can be executed (i.e., the knight is asleep); {@code false} otherwise.
     */
    public static boolean canFastAttack(boolean knightIsAwake) {
        return !knightIsAwake;
    }

    /**
     * Determines if spying is possible based on the wakefulness of at least one character in the scenario.
     *
     * @param knightIsAwake  the wakefulness of the knight.
     * @param archerIsAwake  the wakefulness of the archer.
     * @param prisonerIsAwake the wakefulness of the prisoner.
     * @return {@code true} if spying is possible (i.e., at least one of the characters is awake); {@code false} otherwise.
     */
    public static boolean canSpy(boolean knightIsAwake, boolean archerIsAwake, boolean prisonerIsAwake) {
        return knightIsAwake || archerIsAwake || prisonerIsAwake;
    }

    /**
     * Determines if the prisoner can be signaled based on the prisoner's and the archer's wakefulness.
     *
     * @param archerIsAwake  the wakefulness of the archer.
     * @param prisonerIsAwake the wakefulness of the prisoner.
     * @return {@code true} if the prisoner can be signaled (i.e., the prisoner is awake and the archer is asleep); {@code false} otherwise.
     */
    public static boolean canSignalPrisoner(boolean archerIsAwake, boolean prisonerIsAwake) {
        return prisonerIsAwake && !archerIsAwake;
    }

    /**
     * Determines if freeing the prisoner is possible based on the wakefulness of the knight, the archer, and the prisoner, as well as the presence of Annalyn's pet dog.
     *
     * @param knightIsAwake   the wakefulness of the knight.
     * @param archerIsAwake   the wakefulness of the archer.
     * @param prisonerIsAwake the wakefulness of the prisoner.
     * @param petDogIsPresent {@code true} if Annalyn's pet dog is present; {@code false} otherwise.
     * @return {@code true} if freeing the prisoner is possible; {@code false} otherwise. This is true if either the prisoner is awake and both the knight and archer are asleep, or if Annalyn's pet dog is present and the archer is asleep.
     */
    public static boolean canFreePrisoner(boolean knightIsAwake, boolean archerIsAwake, boolean prisonerIsAwake, boolean petDogIsPresent) {
        return (petDogIsPresent && !archerIsAwake) || (prisonerIsAwake && !knightIsAwake && !archerIsAwake);
    }
}
