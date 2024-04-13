import java.util.Random;

class CaptainsLog {

    private static final char[] PLANET_CLASSES = new char[]{'D', 'H', 'J', 'K', 'L', 'M', 'N', 'R', 'T', 'Y'};

    private Random random;

    CaptainsLog(Random random) {
        this.random = random;
    }

    char randomPlanetClass() {
        var randomIndex = random.nextInt(PLANET_CLASSES.length);
        return PLANET_CLASSES[randomIndex];
    }

    String randomShipRegistryNumber() {
        var registryNumber = 1000 + random.nextInt(9000);
        return "NCC-" + registryNumber;
    }

    double randomStardate() {
        // Star Dates are between 41000.00 and 42000.00 (exclusive)
        return 41000.00 + random.nextDouble(1000);
    }
}
