
class BirdWatcher {
    private final int[] birdsPerDay;

    public BirdWatcher(int[] birdsPerDay) {
        this.birdsPerDay = birdsPerDay.clone();
    }

    public int[] getLastWeek() {
        return new int[] { 0, 2, 5, 3, 7, 8, 4 };
    }

    public int getToday() {
        return birdsPerDay[birdsPerDay.length - 1];
    }

    public void incrementTodaysCount() {
        birdsPerDay[birdsPerDay.length - 1]++;
    }

    public boolean hasDayWithoutBirds() {
        for (var birdCount: birdsPerDay) {
            if (birdCount == 0) {
                return true;
            }
        }

        return false;
    }

    public int getCountForFirstDays(int numberOfDays) {
        var count = 0;

        // Number of days might exceed the length of the birdsPerDay array.
        numberOfDays = Math.min(numberOfDays, birdsPerDay.length);

        for (int i = 0; i < numberOfDays; i++) {
            count += birdsPerDay[i];
        }

        return count;
    }

    public int getBusyDays() {
        var busyDays = 0;

        for (var birdCount: birdsPerDay) {
            if (birdCount >= 5) {
                busyDays++;
            }
        }

        return busyDays;
    }
}
