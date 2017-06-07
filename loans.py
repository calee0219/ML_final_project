import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import subprocess
import shlex

loans = pd.read_csv("./loans_data.csv")
# loans.plot(kind="scatter", x="dti", y="safe_loans")
sns.pairplot(loans, hue="safe_loans", size=3)
sns.plt.show()

fname='./test.pdf'
plt.savefig(fname)
proc=subprocess.Popen(shlex.split('lpr {f}'.format(f=fname)))