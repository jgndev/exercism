public class SalaryCalculator {

    /**
     * The maximum possible salary an employee can earn.
     */
    private static final double SALARY_CAP = 2_000.00;
    /**
     * The base salary before any adjustments for attendance or sales performance.
     */
    private static final double SALARY_BASE = 1_000.00;

    /**
     * Calculates the salary multiplier based on the number of days an employee has skipped work.
     * <p>
     *     A 15% penalty is applied to the salary for employees who have skipped work 5 or more days.
     * </p>
     * @param daysSkipped The number of days the employee has skipped work.
     * @return A multiplier to adjust the base salary. Less than 5 days skipped results in no penalty.
     */
    public double salaryMultiplier(int daysSkipped) {
        // Employees who skipped 5 or more days receive a 15% penalty.
        return daysSkipped >= 5 ? 0.85 : 1.0;
    }

    /**
     * Determines the bonus multiplier based on the number of products an employee has sold.
     * <p>
     * Employees who sell more than 20 products receive a higher per-unit bonus.
     * </p>
     *
     * @param productsSold The total number of products sold by the employee.
     * @return The per-unit bonus amount.
     */
    public int bonusMultiplier(int productsSold) {
        // Employees who sell more than 20 products get 13 per unit sold, less than 20 gets 10.
        return productsSold >= 20 ? 13 : 10;
    }

    /**
     * Calculates the total bonus earned from products sold.
     * <p>
     * The bonus is determined by multiplying the number of products sold by the applicable bonus multiplier.
     * </p>
     *
     * @param productsSold The total number of products sold by the employee.
     * @return The total bonus amount earned from sales.
     */
    public double bonusForProductsSold(int productsSold) {
        // Bonus is the multiplier * the products sold.
        return bonusMultiplier(productsSold) * productsSold;
    }

    /**
     * Calculates the final salary of an employee after adjusting for days skipped and bonuses for products sold.
     * <p>
     * The final salary is the sum of the adjusted base salary and sales bonuses, capped at a maximum value.
     * </p>
     *
     * @param daysSkipped  The number of days the employee skipped work.
     * @param productsSold The total number of products sold by the employee.
     * @return The final calculated salary of the employee, subject to a salary cap.
     */
    public double finalSalary(int daysSkipped, int productsSold) {
        var salaryMultiplier = salaryMultiplier(daysSkipped);
        var bonus = bonusForProductsSold(productsSold);
        var adjustedSalary = (SALARY_BASE * salaryMultiplier) + bonus;

//        return Math.min(adjustedSalary, SALARY_CAP);
        return adjustedSalary >= SALARY_CAP ? SALARY_CAP : adjustedSalary;
    }
}
