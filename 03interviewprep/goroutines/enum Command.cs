/* git add/status/ls/branch
/* git <command> <subcommand> <optional> <parameters>
*/

enum Command
{
    ADD
    STATUS
    BRANCH
}

enum SubCommand
{
    RENAME
    DELETE
}

public abstract class Command
{
    private string name;

    public abstract string GetCommandName();
}

class AddCommand : Command
{
    public override string GetCommandName()
    {
        return "add";
    }
}

/* git add/status/ls/branch
/* git <command> <subcommand> <optional> <parameters>

git branch rename
git remote rename

*/
class StatusCommand : Command
{
     public override string GetCommandName()
    {
        return "status";
    }
}

class BranchCommand : Command
{
     public override string GetCommandName()
    {
        return "branch";
    }
}

public abstract class SubCommand
{
    private string subCommandName;
    public abstract string GetSubCommandName();
}

class Checkout : SubCommand
{


    public override string GetSubCommandName()
    {
        return "checkout";
    }
}

// git branch rename
// git remote rename

class Rename : SubCommand
{
    private HashSet<Command> commandHash = new HashSet<Command>();

    public Rename(List<Command> command)
    {
        for(int i = 0; i < command.Count; i++)
        {
            commandHash.Add(command[i]);
        }
    }

    public override string GetSubCommandName()
    {
        return "rename";
    }

    public bool IsValidCommand(string fullCommand)
    {

    }
}





class Program
{

    public static void Main(string[] args)
    {
        string command = args[1];
        string subCommand = args[2];

        HashSet<string> commandHash = new HashSet<string>();
        Enum.GetValues<Command>().ToList().ForEach(x => commandHash.Add(x.ToString()));
        
        if(!commandHash.Contains(command.ToUpperCase()))
        {
            throw new Error("Invalid Command" + command);
        }

        HashSet<string> subCommandHash = new HashSet<string>();
        Enum.GetValues<SubCommand>().ToList().ForEach(x => subCommandHash.Add(x.ToString()));

        if(!subCommandHash.Contains(subCommand.ToUpperCase()))
        {
            throw new Error("Invalid SubCommand" + subCommand);
        }

        if(subCommand.ToUpperCase() == SubCommand.RENAME)
        {
            
        }
        
    }
}


// git checkout rename - -ve case
// git checkout goodmorning 
// git branch rename
// git remote rename
