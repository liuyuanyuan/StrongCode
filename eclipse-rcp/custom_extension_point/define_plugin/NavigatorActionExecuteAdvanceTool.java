package com.highgo.admin;

import java.util.Collection;
import java.util.List;

import org.eclipse.jface.action.IAction;
import org.eclipse.jface.viewers.ISelection;
import org.eclipse.ui.IActionDelegate;
import org.eclipse.ui.IWorkbenchPart;
import org.eclipse.ui.IWorkbenchWindow;
import org.jkiss.dbeaver.core.DBeaverUI;
import org.jkiss.dbeaver.model.struct.DBSObject;
import org.jkiss.dbeaver.runtime.ui.DBUserInterface;
import org.jkiss.dbeaver.tools.IExternalTool;
import org.jkiss.dbeaver.ui.navigator.NavigatorUtils;

public class NavigatorActionExecuteAdvanceTool  implements IActionDelegate
{
		
    private IWorkbenchWindow window;
    private AdvanceToolDescriptor tool;
    
    private ISelection selection;

    public NavigatorActionExecuteAdvanceTool(IWorkbenchWindow window, AdvanceToolDescriptor tool)
    {
        this.window = window;
        this.tool = tool;
    }

    @Override
    public void run(IAction action)
    {
        if (!selection.isEmpty()) 
        {
            List<DBSObject> selectedObjects = NavigatorUtils.getSelectedObjects(selection);
            executeTool(DBeaverUI.getActiveWorkbenchWindow().getActivePage().getActivePart(), selectedObjects);
        }
    }

    private void executeTool(IWorkbenchPart part, Collection<DBSObject> objects)
    {
        try 
        {
            IExternalTool toolInstance = tool.createTool();
            toolInstance.execute(window, part, objects);
        } catch (Throwable e) {
            DBUserInterface.getInstance().showError("Tool error", "Error executing tool '" + tool.getLabel() + "'", e);
        }
    }

    @Override
    public void selectionChanged(IAction action, ISelection selection)
    {
        this.selection = selection;
    }

}